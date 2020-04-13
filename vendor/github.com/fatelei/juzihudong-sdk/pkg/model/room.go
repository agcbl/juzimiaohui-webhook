package model

type RoomMember struct {
	Wxid string `json:"wxid,omitempty"`
	NickName string `json:"nickName,omitempty"`
	AvatarUrl string `json:"avatarUrl,omitempty"`
	RoomAlias string `json:"roomAlias,omitempty"`
	IsFriend bool `json:"isFriend,omitempty"`
}

type Room struct {
	ChatId string `json:"chatId,omitempty"`
	Members *[]RoomMember `json:"members,omitempty"`
	BotInfo *BotInfo `json:"botInfo,omitempty"`
	WxId string `json:"wxid,omitempty"`
	Topic string `json:"topic,omitempty"`
	AvatarUrl string `json:"avatarUrl,omitempty"`
	OwnerId string `json:"ownerId,omitempty"`
}
