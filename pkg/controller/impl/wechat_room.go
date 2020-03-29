package impl

import (
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao/impl"
	"github.com/fatelei/juzimiaohui-webhook/pkg/model"
)

type WechatRoomControllerImpl struct {}

var DefaultWechatRoomController *WechatRoomControllerImpl

func init() {
	DefaultWechatRoomController = &WechatRoomControllerImpl{}
}

func (p *WechatRoomControllerImpl) CreateRoom(room *model.Room) {
	impl.DefaultWechatRoomDAOImpl.Create(room)
}
