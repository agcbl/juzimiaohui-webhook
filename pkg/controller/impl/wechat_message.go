package impl

import (
	"github.com/fatelei/juzihudong-sdk/pkg/juzihudong"
	"github.com/fatelei/juzimiaohui-webhook/configs"
	"github.com/fatelei/juzimiaohui-webhook/pkg/controller"
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao/impl"
	"github.com/fatelei/juzimiaohui-webhook/pkg/model"
	"log"
)

type WechatMessageControllerImpl struct {
	contactApi *juzihudong.ContactApi
}
var _ controller.WechatMessageController = (*WechatMessageControllerImpl)(nil)


func NewWechatMessageController() *WechatMessageControllerImpl {
	contactApi := juzihudong.NewContactApi(configs.DefaultConfig.Juzihudong.Endpoint, configs.DefaultConfig.Juzihudong.Token)
	wechatMessageController := &WechatMessageControllerImpl{
		contactApi:contactApi,
	}
	return wechatMessageController
}


func (p *WechatMessageControllerImpl) recordActive(message *model.WechatMessage) {
	recordId := impl.DefaultWechatUserInfoDAO.Get(message.ContactId, message.RoomId)
	if recordId > 0 {
		impl.DefaultWechatUserInfoDAO.UpdateLastActiveTime(recordId)
		log.Printf("update last active time %d\n", recordId)
	} else {
		resp := p.contactApi.GetContact(0, 1, message.ContactId)
		if len(*resp.Data) > 0 {
			for _, contact := range *resp.Data {
				impl.DefaultWechatUserInfoDAO.Create(
					contact.Weixin, message.ContactId, message.RoomId, message.ContactName, int(contact.Gender), contact.City, contact.Province, contact.AvatarUrl)
				log.Printf("record active with %+v\n", contact)
				break
			}
		} else {
			impl.DefaultWechatUserInfoDAO.Create(
				"", message.ContactId, message.RoomId, message.ContactName, 0, "", "", "")
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
		DefaultNotificationController.CreateNotification(
			wechatMessage.RoomTopic, wechatMessage.ContactName, wechatMessage.GetContent())
		p.recordActive(wechatMessage)
	} else {
		log.Printf("not support group %s\n", wechatMessage.RoomId)
	}
}
