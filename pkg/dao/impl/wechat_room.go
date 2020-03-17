package impl

import (
	"database/sql"
	"github.com/fatelei/juzihudong-sdk/pkg/model"
	"github.com/fatelei/juzimiaohui-webhook/pkg/connection"
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao"
	model2 "github.com/fatelei/juzimiaohui-webhook/pkg/model"
	"log"
)

type WechatRoomDAOImpl struct {}

var DefaultWechatRoomDAOImpl *WechatRoomDAOImpl
var _ dao.WechatRoomDAO = (*WechatRoomDAOImpl)(nil)

func init() {
	DefaultWechatRoomDAOImpl = &WechatRoomDAOImpl{}
}

func (p *WechatRoomDAOImpl) Create(room *model.Room) {
	var result sql.Result
	stmtIns, err := connection.DB.Prepare(
		"INSERT INTO wechat_room (room_id, room_name) VALUES(?, ?)")
	if err != nil {
		panic(err)
	}

	defer stmtIns.Close()
	result, err = stmtIns.Exec(room.WxId, room.Topic)
	if err != nil {
		panic(err)
	} else {
		_id, _ := result.LastInsertId()
		log.Printf("insert success: %d\n", _id)
	}
}

func (p *WechatRoomDAOImpl) GetRoomByRoomId(roomId string) *model2.Room {
	stmtQuery, err := connection.DB.Prepare(
		"SELECT * FROM wechat_room WHERE room_id = ?")
	if err != nil {
		panic(err)
	}

	defer stmtQuery.Close()
	room := &model2.Room{}
	row := stmtQuery.QueryRow(roomId)
	if row != nil {
		err := row.Scan(&room.Id, &room.RoomId, &room.RoomName, &room.RoomMemberNumber, &room.OpenMonitor)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			} else {
				panic(err)
			}
		}
		return room
	}
	return nil
}
