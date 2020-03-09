package dao


type WechatUserInfoDAO interface {
	Create(wechatId string, wxid string, wechatName string, gender int, city string, province string, avatarUrl string)
	Get(wxid string) int64
	UpdateLastActiveTime(recordId int64)
}
