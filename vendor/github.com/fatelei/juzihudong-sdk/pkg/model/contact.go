package model

type Contact struct {
	ChatId string `json:"chatId,omitempty"`
	BotInfo *BotInfo `json:"botInfo,omitempty"`
	WxId string `json:"wxid,omitempty"`
	Weixin string `json:"weixin,omitempty"`
	Nickname string `json:"nickName,omitempty"`
	Alias string `json:"alias,omitempty"`
	AvatarUrl string `json:"avatarUrl,omitempty"`
	City string `json:"city,omitempty"`
	Province string `json:"province,omitempty"`
	Gender int64 `json:"gender,omitempty"`
}
