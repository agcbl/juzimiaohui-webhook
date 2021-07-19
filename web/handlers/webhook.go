package handlers

import (
	"github.com/fatelei/juzimiaohui-webhook/configs"
	impl "github.com/fatelei/juzimiaohui-webhook/pkg/controller/impl"
	impl2 "github.com/fatelei/juzimiaohui-webhook/pkg/dao/impl"
	"github.com/fatelei/juzimiaohui-webhook/pkg/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type WebhookHandler struct {
	wechatMessageController *impl.WechatMessageControllerImpl
}

type qrcode struct {
	URL string `json:"url"`
}

func NewWebhookHandler() *WebhookHandler {
	wechatMessageController := impl.NewWechatMessageController()
	return &WebhookHandler{wechatMessageController:wechatMessageController}
}

func (p *WebhookHandler) QRCodeCallback(c *gin.Context) {
	var data qrcode
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(data.URL)
	p.wechatMessageController.SendLoginUrl(data.URL)
	c.JSON(http.StatusCreated, gin.H{})
	return
}


func (p *WebhookHandler) MessageCallback(c *gin.Context) {
	var data model.WechatMessageData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("room %s received message", data.Data.RoomTopic)
	isInWhiteList := impl2.DefaultWechatWhiteListDAO.IsInWhiteList(data.Data.ContactId)
	if !isInWhiteList {
		if data.Data.RoomId == configs.DefaultConfig.WhiteList.RoomID {
			impl2.DefaultWechatWhiteListDAO.AddWhiteListMember(data.Data.ContactId)
		} else {
			p.wechatMessageController.Create(&data.Data)
		}
	}
	c.JSON(http.StatusCreated, gin.H{})
	return
}

