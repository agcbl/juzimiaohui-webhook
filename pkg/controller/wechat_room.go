package controller

import "github.com/fatelei/juzihudong-sdk/pkg/model"

type WechatRoomController interface {
	CreatRoom(room *model.Room)
}
