package controller

type NotificationController interface {
	CreateNotification(room string, contactName string, contentId string, content string)
	CreateWechatDeathNoti()
}
