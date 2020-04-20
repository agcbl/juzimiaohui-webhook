package impl

import (
	"github.com/fatelei/juzimiaohui-webhook/pkg/connection"
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao"
)

type WechatWhiteListDAOImpl struct {}

var DefaultWechatWhiteListDAO dao.WhiteListDAO
var _ dao.WhiteListDAO = (*WechatWhiteListDAOImpl)(nil)

func init() {
	DefaultWechatWhiteListDAO = NewWechatWhiteListDAO()
}

func NewWechatWhiteListDAO() *WechatWhiteListDAOImpl {
	return &WechatWhiteListDAOImpl{}
}

func (p *WechatWhiteListDAOImpl) IsInWhiteList(wxid string) bool {
	stmtQuery, err := connection.DB.Prepare(
		"SELECT id FROM wechat_whitelist WHERE wxid = ?")
	if err != nil {
		panic(err)
	}

	defer stmtQuery.Close()
	row := stmtQuery.QueryRow(wxid)
	if row != nil {
		var _id int64
		err := row.Scan(&_id)
		if err != nil {
			return false
		}
		return _id > 0
	}
	return false
}

func (p *WechatWhiteListDAOImpl) AddWhiteListMember(wxid string) {
	stmtIns, err := connection.DB.Prepare(
		"INSERT INTO wechat_whitelist (wxid) VALUES (?)")
	if err != nil {
		panic(err)
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(wxid)
	if err != nil {
		panic(err)
	}
}