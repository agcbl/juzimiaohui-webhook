package controller

import "github.com/fatelei/juzimiaohui-webhook/pkg/model"

type WechatMessageController interface {
	Create(wechatMessage *model.WechatMessage)
}
