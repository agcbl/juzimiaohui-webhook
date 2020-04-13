package model

type BotInfo struct {
	Wxid string `json:"wxid,omitempty"`
	Weixin string `json:"weixin,omitempty"`
	Nickname string `json:"nickName,omitempty"`
}
