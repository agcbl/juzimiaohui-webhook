package impl

import (
	"github.com/fatelei/juzimiaohui-webhook/pkg/connection"
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao"
)

type WechatRoomMemberDAOImpl struct {}

var _ dao.WechatRoomMemberDAO = (*WechatRoomMemberDAOImpl)(nil)
var DefaultWechatRoomMemberDAO dao.WechatRoomMemberDAO

func init() {
	DefaultWechatRoomMemberDAO = NewWechatRoomMemberDAO()
}

func NewWechatRoomMemberDAO() *WechatRoomMemberDAOImpl {
	return &WechatRoomMemberDAOImpl{}
}


func (p *WechatRoomMemberDAOImpl) GetRoomAlias(roomID string, wxIDs []string) map[string]string {
	stmtQuery, err := connection.DB.Prepare(
		"SELECT wxid, room_alias FROM wechat_room_member WHERE room_id = ? AND wxid in (?)")
	if err != nil {
		panic(err)
	}

	defer stmtQuery.Close()
	result := make(map[string]string)
	rows, _ := stmtQuery.Query(roomID, wxIDs)
	for rows.Next() {
		var wxid string
		var roomAlias string
		err := rows.Scan(&wxid, &roomAlias)
		if err == nil {
			result[wxid] = roomAlias
		}
	}
	return result
}
