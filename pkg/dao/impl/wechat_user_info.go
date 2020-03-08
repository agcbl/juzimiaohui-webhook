package impl

import (
	"github.com/fatelei/juzimiaohui-webhook/pkg/connection"
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao"
	"time"
)

type WechatUserInfoDAOImpl struct {}

var DefaultWechatUserInfoDAO *WechatUserInfoDAOImpl
var _ dao.WechatUserInfoDAO = (*WechatUserInfoDAOImpl)(nil)

func init() {
	DefaultWechatUserInfoDAO = &WechatUserInfoDAOImpl{}
}


func (p *WechatUserInfoDAOImpl) Create(wechatId string, wxid string, roomId string, wechatName string, gender int, city string, province string, avatarUrl string) {
	stmtIns, err := connection.DB.Prepare(
		"INSERT INTO wechat_user_info (wechat_id, room_id, wxid, wechat_name, gender, city, province, avatar_url) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}

	defer stmtIns.Close()
	_, err = stmtIns.Exec(wechatId, roomId, wxid, wechatName, gender, city, province, avatarUrl)
	if err != nil {
		panic(err)
	}
}

func (p *WechatUserInfoDAOImpl) Get(wxid string, roomId string) int64 {
	stmtQuery, err := connection.DB.Prepare(
		"SELECT id FROM wechat_user_info WHERE wxid = ? and room_id = ?")
	if err != nil {
		panic(err)
	}

	defer stmtQuery.Close()
	var id int64
	row := stmtQuery.QueryRow(wxid, roomId)
	if row != nil {
		err := row.Scan(&id)
		if err != nil {
			return 0
		}
		return id
	}
	return 0
}


func (p *WechatUserInfoDAOImpl) UpdateLastActiveTime(recordId int64) {
	stmtUpdate, err := connection.DB.Prepare(
		"UPDATE wechat_user_info SET last_active_time = ? WHERE id = ?")
	if err != nil {
		panic(err)
	}

	local, _ := time.LoadLocation("Asia/Chongqing")
	defer stmtUpdate.Close()
	_, err = stmtUpdate.Exec(time.Now().In(local).Format(time.RFC3339), recordId)
	if err != nil {
		panic(err)
	}
}
