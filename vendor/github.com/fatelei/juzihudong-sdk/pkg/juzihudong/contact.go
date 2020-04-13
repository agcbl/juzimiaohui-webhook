package juzihudong

import (
	"encoding/json"
	"github.com/fatelei/juzihudong-sdk/pkg/model"
	"github.com/fatelei/juzihudong-sdk/pkg/transport"
	"strconv"
)

type ContactApi struct {
	Transport *transport.Transport
}

type ContactListResponse struct {
	Code int64 `json:"code,omitempty"`
	Data *[]model.Contact `json:"data,omitempty"`
	Page *model.Page `json:"page,omitempty"`
}


func NewContactApi(endpoint string, token string) *ContactApi {
	transport := &transport.Transport{Endpoint:endpoint,Token:token,}
	contactApi := &ContactApi{Transport:transport}
	return contactApi
}


func (p *ContactApi) GetContact(current int, pageSize int, wxid string) *ContactListResponse {
	param := make(map[string]string)
	if len(wxid) > 0 {
		param["wxid"] = wxid
	} else {
		param["current"] = strconv.Itoa(current)
		param["pageSize"] = strconv.Itoa(pageSize)
	}
	body, err := p.Transport.Get("/contact/list", &param)
	if err != nil {
		panic(err)
	}
	resp := ContactListResponse{}
	if err := json.Unmarshal(body, &resp); err == nil {
		return &resp
	}
	return nil
}
