package impl

import (
	"github.com/fatelei/juzihudong-sdk/pkg/juzihudong"
	"github.com/fatelei/juzimiaohui-webhook/configs"
	"github.com/fatelei/juzimiaohui-webhook/pkg/controller"
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao/impl"
	"github.com/fatelei/juzimiaohui-webhook/pkg/model"
	"log"
	"strconv"
	"time"
)

type WechatMessageControllerImpl struct {
	contactApi *juzihudong.ContactApi
	meesageApi *juzihudong.MessageApi
	notificationController *NotificationControllerImpl
	recentMessageId int64
	duplicateCount int
}
var _ controller.WechatMessageController = (*WechatMessageControllerImpl)(nil)

func NewWechatMessageController() *WechatMessageControllerImpl {
	contactApi := juzihudong.NewContactApi(configs.DefaultConfig.Juzihudong.Endpoint, configs.DefaultConfig.Juzihudong.Token)
	meesageApi := juzihudong.NewMessageApi(configs.DefaultConfig.Juzihudong.Endpoint, configs.DefaultConfig.Juzihudong.Token)
	notificationController := NewNotificationController()
	recentMessageId := impl.DefaultWechatMessageDAO.GetMaxMessageId()
	wechatMessageController := &WechatMessageControllerImpl{
		contactApi:contactApi,
		meesageApi:meesageApi,
		notificationController: notificationController,
		recentMessageId: recentMessageId,
		duplicateCount: 0,
	}
	if configs.DefaultConfig.Alive != nil {
		wechatMessageController.checkAlive()
	}
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
		log.Println("not group message")
		return
	}

	var flag = false
	room := impl.DefaultWechatRoomDAOImpl.GetRoomByRoomId(wechatMessage.RoomId)
	if room == nil {
		flag = true
		if wechatMessage.Type != 10001 {
			tmp := model.Room{
				Id:               0,
				RoomId:           wechatMessage.RoomId,
				RoomName:         wechatMessage.RoomTopic,
				RoomMemberNumber: 0,
				OpenMonitor:      0,
			}
			impl.DefaultWechatRoomDAOImpl.Create(&tmp)
			log.Printf("create new room %+v\n", tmp)
		}
	} else if room.OpenMonitor == 1 {
		flag = true
	} else {
		log.Printf("not support group %s\n", wechatMessage.RoomId)
	}

	if flag {
		if wechatMessage.Type == model.Image {
			image := p.meesageApi.GetArtworkImage(wechatMessage.ChatId, wechatMessage.MessageId)
			if image != nil && image.Code == 0 && image.Data != nil {
				wechatMessage.Payload.ImageUrl = image.Data.Url
			}
		}
		impl.DefaultWechatMessageDAO.Create(wechatMessage)
		if wechatMessage.Type == model.Image {
			go p.notificationController.CreateMessageCard(wechatMessage)
		} else {
			p.notificationController.CreateNotification(wechatMessage)
		}
		p.recordActive(wechatMessage)
	}
}

func (p *WechatMessageControllerImpl) GetRecentMessages(
	chatID string, wxid string, roomId string, createdAt string, direction string, action string) {
	timestamp, _ := strconv.Atoi(createdAt)
	tm := time.Unix(int64(timestamp/1000), 0)
	createdAtStr := tm.Format("2006-01-02 15:04:05")
	messages := impl.DefaultWechatMessageDAO.GetRecentMessages(wxid, roomId, createdAtStr, direction)
	if len(messages) > 0 {
		if action == "loadMyRoomMessage" {
			p.notificationController.SendRecentMessagesCard(chatID, messages)
		} else if action == "loadRoomMessage" {
			wxids := make([]string, len(messages))
			roomAlias := impl.DefaultWechatRoomMemberDAO.GetRoomAlias(roomId, wxids)
			p.notificationController.SendRoomRecentMessagesCard(chatID, messages, roomAlias)
		}
	}
}


func (p *WechatMessageControllerImpl) checkAlive() {
	go func() {
		for {
			time.Sleep(time.Duration(configs.DefaultConfig.Alive.Tick) * time.Minute)
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