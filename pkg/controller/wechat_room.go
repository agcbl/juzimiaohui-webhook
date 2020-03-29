package controller

import (
	"github.com/fatelei/juzimiaohui-webhook/pkg/model"
)


type WechatRoomController interface {
	CreatRoom(room *model.Room)
}
