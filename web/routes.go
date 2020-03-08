package web

import (
	"github.com/fatelei/juzimiaohui-webhook/web/handlers"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	webhookHandler := handlers.NewWebhookHandler()
	r.POST("/message", webhookHandler.MessageCallback)
	return r
}
