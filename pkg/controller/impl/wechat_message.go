package impl

import (
	"github.com/fatelei/juzimiaohui-webhook/pkg/controller"
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao/impl"
	"github.com/fatelei/juzimiaohui-webhook/pkg/model"
)

type WechatMessageControllerImpl struct {}
var DefaultWechatMessageController *WechatMessageControllerImpl

var _ controller.WechatMessageController = (*WechatMessageControllerImpl)(nil)

func init() {
	DefaultWechatMessageController = &WechatMessageControllerImpl{}
}

func (p *WechatMessageControllerImpl) Create(wechatMessage *model.WechatMessage) {
	impl.DefaultWechatMessageDAO.Create(wechatMessage)
}