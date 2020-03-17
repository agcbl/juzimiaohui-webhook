package handlers

import (
	impl "github.com/fatelei/juzimiaohui-webhook/pkg/controller/impl"
	"github.com/fatelei/juzimiaohui-webhook/pkg/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type WebhookHandler struct {
	wechatMessageController *impl.WechatMessageControllerImpl
}

func NewWebhookHandler() *WebhookHandler {
	wechatMessageController := impl.NewWechatMessageController()
	return &WebhookHandler{wechatMessageController:wechatMessageController}
}


func (p *WebhookHandler) MessageCallback(c *gin.Context) {
	var data model.WechatMessageData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("%+v\n", data.Data)
	p.wechatMessageController.Create(&data.Data)
	c.JSON(http.StatusCreated, gin.H{})
	return
}

