package dao

import (
	model2 "github.com/fatelei/juzihudong-sdk/pkg/model"
	"github.com/fatelei/juzimiaohui-webhook/pkg/model"
)


type WechatRoomDAO interface {
	Create(room *model2.Room)
	GetRoomByRoomId(roomId string) *model.Room
}
