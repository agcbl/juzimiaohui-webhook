package impl

import (
	"database/sql"
	"github.com/fatelei/juzimiaohui-webhook/pkg/connection"
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao"
)

type FeishuBotDAOImpl struct {}

var DefaultFeishuBotDAO dao.FeishuBotDAO

var _ dao.FeishuBotDAO = (*FeishuBotDAOImpl)(nil)

func init() {
	DefaultFeishuBotDAO = &FeishuBotDAOImpl{}
}

func (p *FeishuBotDAOImpl) GetAccessToken() *dao.FeishuBotRecord {
	stmtQuery, err := connection.DB.Prepare(
		"SELECT id, access_token, expire FROM feishu_bot")
	if err != nil {
		panic(err)
	}

	defer stmtQuery.Close()
	record := &dao.FeishuBotRecord{}
	row := stmtQuery.QueryRow()
	if row != nil {
		err := row.Scan(&record.ID, &record.Token, &record.Expire)
		if err != nil {
			if err != sql.ErrNoRows {
				panic(err)
			}
			return nil
		}
		return record
	}
	return nil
}

func (p *FeishuBotDAOImpl) Create(token string, expire int64) {
	stmtIns, err := connection.DB.Prepare(
		"INSERT INTO feishu_bot (access_token, expire) VALUES(?, ?)")
	if err != nil {
		panic(err)
	}

	defer stmtIns.Close()
	_, err = stmtIns.Exec(token, expire)
	if err != nil {
		panic(err)
	}
}

func (p *FeishuBotDAOImpl) Refresh(id int64, token string, expire int64) {
	stmtUpdate, err := connection.DB.Prepare(
		"UPDATE feishu_bot SET access_token = ?, expire = ? WHERE id = ?")
	if err != nil {
		panic(err)
	}

	defer stmtUpdate.Close()
	_, err = stmtUpdate.Exec(token, expire, id)
	if err != nil {
		panic(err)
	}
}