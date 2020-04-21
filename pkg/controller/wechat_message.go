package controller

import "github.com/fatelei/juzimiaohui-webhook/pkg/model"

type WechatMessageController interface {
	Create(wechatMessage *model.WechatMessage)
	GetRecentMessages(chatID string, wxid string, roomId string, createdAt string, direction string)
}
