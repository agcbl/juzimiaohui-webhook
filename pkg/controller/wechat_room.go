package controller

type WechatRoomController interface {
	ListRooms(page int64, pageSize int64)
}
