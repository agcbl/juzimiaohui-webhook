package dao

import "github.com/fatelei/juzimiaohui-webhook/pkg/model"

type WechatMessageDAO interface {
	Create(wechatMessage *model.WechatMessage)
	GetMaxMessageId() int64
}
