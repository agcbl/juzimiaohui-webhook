package impl

import (
	"github.com/fatelei/juzimiaohui-webhook/pkg/connection"
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao"
	"github.com/fatelei/juzimiaohui-webhook/pkg/model"
)

type KeywordDAOImpl struct {}

var DefaultKeywordDAO *KeywordDAOImpl

var _ dao.KeywordDAO = (*KeywordDAOImpl)(nil)

func init() {
	DefaultKeywordDAO = &KeywordDAOImpl{}
}

func (p *KeywordDAOImpl) GetKeywords() []model.Keyword {
	stmtQuery, err := connection.DB.Prepare("SELECT * FROM wechat_keywords WHERE is_opened = 1")
	if err != nil {
		panic(err)
	}

	defer stmtQuery.Close()
	words := make([]model.Keyword, 0)
	rows, err := stmtQuery.Query()
	if err != nil {
		return words
	} else {
		for rows.Next() {
			tmp := model.Keyword{}
			rows.Scan(&tmp.Id, &tmp.Word, &tmp.IsOpened)
			words = append(words, tmp)
		}
		return words
	}
}
