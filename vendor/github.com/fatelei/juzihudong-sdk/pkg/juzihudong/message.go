package juzihudong

import (
	"encoding/json"
	"github.com/fatelei/juzihudong-sdk/pkg/model"
	"github.com/fatelei/juzihudong-sdk/pkg/transport"
)

type MessageApi struct {
	Transport *transport.Transport
}

type GetImageResponse struct {
	Code int64 `json:"code,omitempty"`
	Data *model.Image `json:"data,omitempty"`
}


func NewMessageApi(endpoint string, token string) *MessageApi {
	transport := &transport.Transport{Endpoint:endpoint,Token:token}
	messageApi := &MessageApi{Transport:transport}
	return messageApi
}


func (p *MessageApi) SendTextMessage(chatId string, content string) bool {
	body := make(map[string]interface{})
	body["chatId"] = chatId
	body["messageType"] = 1
	body["payload"] = map[string]string{"text": content}
	_, err := p.Transport.Post("/message/send", body)
	if err != nil {
		return false
	}
	return true
}

func (p *MessageApi) GetArtworkImage(chatId string, messageId string) *GetImageResponse {
	param := make(map[string]interface{})
	param["chatId"] = chatId
	param["messageId"] = messageId
	body, err := p.Transport.Post("/message/getArtworkImage", param)
	if err != nil {
		return nil
	}
	resp := GetImageResponse{}
	if err := json.Unmarshal(body, &resp); err == nil {
		return &resp
	}
	return nil
}