package impl

import (
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
		"INSERT INTO wechat_message_monitor (wechat_id, wechat_name, room_name, content, msg_type, room_id) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}

	defer stmtIns.Close()
	_, err = stmtIns.Exec(wechatMessage.ContactId, wechatMessage.ContactName,
		wechatMessage.RoomTopic, wechatMessage.GetContent(), wechatMessage.Type, wechatMessage.RoomId)
	if err != nil {
		panic(err)
	}
}
