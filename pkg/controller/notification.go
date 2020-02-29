package controller

type NotificationController interface {
	CreateNotification(room string, contact string, content string)
}
