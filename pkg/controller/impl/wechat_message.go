package impl

import (
	"github.com/fatelei/juzimiaohui-webhook/pkg/controller"
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao/impl"
	"github.com/fatelei/juzimiaohui-webhook/pkg/model"
	"log"
)

type WechatMessageControllerImpl struct {}
var DefaultWechatMessageController *WechatMessageControllerImpl

var _ controller.WechatMessageController = (*WechatMessageControllerImpl)(nil)

func init() {
	DefaultWechatMessageController = &WechatMessageControllerImpl{}
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
		return
	} else {
		log.Printf("not support group %s\n", wechatMessage.RoomId)
	}
}
