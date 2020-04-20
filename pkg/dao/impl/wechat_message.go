package impl

import (
	"database/sql"
	"github.com/fatelei/juzimiaohui-webhook/pkg/connection"
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao"
	"github.com/fatelei/juzimiaohui-webhook/pkg/model"
)

type WechatMessageDAOImpl struct {}

var _ dao.WechatMessageDAO = (*WechatMessageDAOImpl)(nil)

var DefaultWechatMessageDAO *WechatMessageDAOImpl

func init() {
	DefaultWechatMessageDAO = &WechatMessageDAOImpl{}
}

func (p *WechatMessageDAOImpl) Create(wechatMessage *model.WechatMessage) {
	stmtIns, err := connection.DB.Prepare(
		"INSERT INTO wechat_message_monitor (wxid, wechat_name, room_name, content, msg_type, room_id, message_id) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}

	defer stmtIns.Close()
	_, err = stmtIns.Exec(wechatMessage.ContactId, wechatMessage.ContactName,
		wechatMessage.RoomTopic, wechatMessage.GetContent(), wechatMessage.Type, wechatMessage.RoomId, wechatMessage.MessageId)
	if err != nil {
		panic(err)
	}
}


func (p *WechatMessageDAOImpl) GetMaxMessageId() int64 {
	stmtQuery, err := connection.DB.Prepare(
		"SELECT MAX(id) FROM wechat_message_monitor")
	if err != nil {
		panic(err)
	}

	defer stmtQuery.Close()
	row := stmtQuery.QueryRow()
	var maxId int64
	if row != nil {
		row.Scan(&maxId)
	}
	return maxId
}

func (p *WechatMessageDAOImpl) GetRecentMessages(wxid string, roomId string, createdAt string, direction string) []*dao.WechatMessage {
	var stmtQuery *sql.Stmt
	var err error
	if direction == "before" {
		stmtQuery, err = connection.DB.Prepare(
			`SELECT id, wxid, wechat_name, room_name, content, msg_type, created_at, room_id, message_id FROM wechat_message_monitor
	WHERE wxid = ? AND room_id = ? AND created_at <= ? order by id asc limit 10`)
	} else {
		stmtQuery, err = connection.DB.Prepare(
			`SELECT id, wxid, wechat_name, room_name, content, msg_type, created_at, room_id, message_id FROM wechat_message_monitor
	WHERE wxid = ? AND room_id = ? AND created_at >= ? order by id asc limit 10`)
	}

	if err != nil {
		panic(err)
	}

	defer stmtQuery.Close()
	results := make([]*dao.WechatMessage, 0)
	rows, err := stmtQuery.Query(wxid, roomId, createdAt)
	if err != nil {
		return results
	}
	for rows.Next() {
		tmp := &dao.WechatMessage{}
		rows.Scan(&tmp.ID, &tmp.WxID, &tmp.WechatName, &tmp.RoomName, &tmp.Content, &tmp.MsgType, &tmp.CreatedAt, &tmp.RoomID, &tmp.MessageID)
		results = append(results, tmp)
	}
	return results
}
