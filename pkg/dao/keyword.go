package dao

import "github.com/fatelei/juzimiaohui-webhook/pkg/model"

type KeywordDAO interface {
	GetKeywords() []model.Keyword
}