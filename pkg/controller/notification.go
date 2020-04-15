package controller

import "github.com/fatelei/juzimiaohui-webhook/pkg/dao"

type NotificationController interface {
	CreateNotification(room string, contactName string, contentId string, content string)
	CreateWechatDeathNoti()
	SendRecentMessagesCard(messages []*dao.WechatMessage)
}
