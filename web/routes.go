package web

import (
	"github.com/fatelei/juzimiaohui-webhook/web/handlers"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	webhookHandler := handlers.NewWebhookHandler()
	feishuHandler := handlers.NewFeishuCallback()
	r.POST("/message", webhookHandler.MessageCallback)
	r.POST("/qr_code", webhookHandler.QRCodeCallback)
	r.POST("/feishu/message_card_interactive", feishuHandler.Callback)
	r.POST("/simple_message", webhookHandler.SinpleMessage)
	return r
}
