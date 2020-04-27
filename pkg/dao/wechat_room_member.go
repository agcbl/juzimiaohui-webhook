package dao

type WechatRoomMemberDAO interface {
	GetRoomAlias(roomID string, wxIDs []string) map[string]string
}
