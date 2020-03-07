package impl

import (
	"github.com/fatelei/juzihudong-sdk/pkg/model"
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao/impl"
)

type WechatRoomControllerImpl struct {}

var DefaultWechatRoomController *WechatRoomControllerImpl

func init() {
	DefaultWechatRoomController = &WechatRoomControllerImpl{}
}

func (p *WechatRoomControllerImpl) CreateRoom(room *model.Room) {
	impl.DefaultWechatRoomDAOImpl.Create(room)
}
