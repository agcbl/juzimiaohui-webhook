package web

import (
	"github.com/fatelei/juzimiaohui-webhook/web/handlers"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	r.POST("/message", handlers.MessageCallback)
	return r
}
