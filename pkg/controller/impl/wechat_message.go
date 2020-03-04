package impl

import (
	"github.com/fatelei/juzimiaohui-webhook/configs"
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

	for _, group := range configs.DefaultConfig.Group.Groups {
		if wechatMessage.RoomId == group {
			log.Printf("receive message: %v\n", wechatMessage)
			impl.DefaultWechatMessageDAO.Create(wechatMessage)
			DefaultNotificationController.CreateNotification(
				wechatMessage.RoomTopic, wechatMessage.ContactName, wechatMessage.GetContent())
			return
		}
	}
}
