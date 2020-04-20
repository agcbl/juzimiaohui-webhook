package dao

type WhiteListDAO interface {
	IsInWhiteList(wxid string) bool
	AddWhiteListMember(wxid string)
}
