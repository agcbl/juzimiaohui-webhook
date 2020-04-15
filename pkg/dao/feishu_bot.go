package dao

import "time"

type FeishuBotRecord struct {
	ID int64
	Token string
	Expire int64
	CreatedAt time.Time
}

type FeishuBotDAO interface {
	GetAccessToken() *FeishuBotRecord
	Create(token string, expire int64)
	Refresh(id int64, token string, expire int64)
}