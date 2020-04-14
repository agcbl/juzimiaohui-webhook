package dao

type FeishuBotRecord struct {
	ID int64
	Token string
	Expire int64
}

type FeishuBotDAO interface {
	GetAccessToken() *FeishuBotRecord
	Create(token string, expire int64)
	Refresh(id int64, token string, expire int64)
}