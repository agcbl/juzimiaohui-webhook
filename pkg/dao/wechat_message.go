package dao

import "github.com/fatelei/juzimiaohui-webhook/pkg/model"

import "time"

type WechatMessage struct {
	ID int64
	WxID string
	WechatName string
	RoomName string
	Content string
	MsgType int64
	CreatedAt time.Time
	RoomID string
	MessageID string
}


type WechatMessageDAO interface {
	Create(wechatMessage *model.WechatMessage)
	GetMaxMessageId() int64
	GetRecentMessages(wxid string, roomId string, createdAt string, direction string) []*WechatMessage
}
