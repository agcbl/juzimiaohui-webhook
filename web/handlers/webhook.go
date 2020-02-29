package handlers

import (
	impl "github.com/fatelei/juzimiaohui-webhook/pkg/controller/impl"
	"github.com/fatelei/juzimiaohui-webhook/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)


func MessageCallback(c *gin.Context) {
	var data model.WechatMessageData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	impl.DefaultWechatMessageController.Create(&data.Data)
	c.JSON(http.StatusCreated, gin.H{})
	return
}

