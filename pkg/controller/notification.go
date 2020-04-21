package controller

import (
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao"
	"github.com/fatelei/juzimiaohui-webhook/pkg/model"
)

type NotificationController interface {
	CreateNotification(message *model.WechatMessage)
	CreateWechatDeathNoti()
	SendRecentMessagesCard(chatID string, messages []*dao.WechatMessage)
}
