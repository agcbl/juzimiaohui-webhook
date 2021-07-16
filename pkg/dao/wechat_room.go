package dao

import (
	"github.com/fatelei/juzimiaohui-webhook/pkg/model"
)


type WechatRoomDAO interface {
	Create(room *model.Room)
	GetRoomByRoomId(roomId string) *model.Room
	GetRoomByRoomTopic(roomTopic string) *model.Room
}
