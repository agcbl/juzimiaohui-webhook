package impl

import (
	"github.com/fatelei/juzihudong-sdk/pkg/juzihudong"
	"github.com/fatelei/juzimiaohui-webhook/configs"
	"github.com/fatelei/juzimiaohui-webhook/pkg/controller"
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao/impl"
	"github.com/fatelei/juzimiaohui-webhook/pkg/model"
	"log"
	"time"
)

type WechatMessageControllerImpl struct {
	contactApi *juzihudong.ContactApi
	notificationController *NotificationControllerImpl
	recentMessageId int64
	duplicateCount int
}
var _ controller.WechatMessageController = (*WechatMessageControllerImpl)(nil)


func NewWechatMessageController() *WechatMessageControllerImpl {
	contactApi := juzihudong.NewContactApi(configs.DefaultConfig.Juzihudong.Endpoint, configs.DefaultConfig.Juzihudong.Token)
	notificationController := NewNotificationController()
	recentMessageId := impl.DefaultWechatMessageDAO.GetMaxMessageId()
	wechatMessageController := &WechatMessageControllerImpl{
		contactApi:contactApi,
		notificationController: notificationController,
		recentMessageId: recentMessageId,
		duplicateCount: 0,
	}
	wechatMessageController.checkAlive()
	return wechatMessageController
}


func (p *WechatMessageControllerImpl) recordActive(message *model.WechatMessage) {
	recordId := impl.DefaultWechatUserInfoDAO.Get(message.ContactId)
	if recordId > 0 {
		impl.DefaultWechatUserInfoDAO.UpdateLastActiveTime(recordId)
		log.Printf("update last active time %d\n", recordId)
	} else {
		resp := p.contactApi.GetContact(0, 1, message.ContactId)
		if len(*resp.Data) > 0 {
			for _, contact := range *resp.Data {
				impl.DefaultWechatUserInfoDAO.Create(
					contact.Weixin, message.ContactId, message.ContactName, int(contact.Gender), contact.City, contact.Province, contact.AvatarUrl)
				log.Printf("record active with %+v\n", contact)
				break
			}
		} else {
			impl.DefaultWechatUserInfoDAO.Create(
				"", message.ContactId, message.ContactName, 0, "", "", "")
			log.Println("record active without wechat_id")
		}
	}
}


func (p *WechatMessageControllerImpl) Create(wechatMessage *model.WechatMessage) {
	if len(wechatMessage.RoomId) == 0 {
		log.Print("not group message")
		return
	}

	room := impl.DefaultWechatRoomDAOImpl.GetRoomByRoomId(wechatMessage.RoomId)
	if room != nil && room.OpenMonitor == 1 {
		log.Printf("receive message: %+v\n", wechatMessage)
		impl.DefaultWechatMessageDAO.Create(wechatMessage)
		p.notificationController.CreateNotification(
			wechatMessage.RoomTopic, wechatMessage.ContactName, wechatMessage.ContactId, wechatMessage.GetContent())
		p.recordActive(wechatMessage)
	} else {
		log.Printf("not support group %s\n", wechatMessage.RoomId)
	}
}


func (p *WechatMessageControllerImpl) checkAlive() {
	go func() {
		for {
			time.Sleep(time.Duration(configs.DefaultConfig.Alive.Tick) * time.Minute)
			hour := time.Now().Hour()
			if hour < configs.DefaultConfig.Alive.StartAt || hour > configs.DefaultConfig.Alive.EndAt {
				log.Printf("hour %d not in change range %d - %d", hour, configs.DefaultConfig.Alive.StartAt, configs.DefaultConfig.Alive.EndAt)
				continue
			}
			currentMaxMessageId := impl.DefaultWechatMessageDAO.GetMaxMessageId()
			log.Printf("check wechat is alive: recent message id %d, current message id %d\n", p.recentMessageId, currentMaxMessageId)
			if currentMaxMessageId == p.recentMessageId {
				p.duplicateCount += 1
				if p.duplicateCount > configs.DefaultConfig.Alive.Limit {
					p.notificationController.CreateWechatDeathNoti()
					p.duplicateCount = 0
				}
			} else {
				p.recentMessageId = currentMaxMessageId
				p.duplicateCount = 0
			}
		}
	}()
}