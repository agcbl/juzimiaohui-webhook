package handlers

import (
	"github.com/fatelei/juzimiaohui-webhook/configs"
	"github.com/fatelei/juzimiaohui-webhook/pkg/controller"
	"github.com/fatelei/juzimiaohui-webhook/pkg/controller/impl"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type FeishuCallback struct {
	wechatMessageController controller.WechatMessageController
}

func NewFeishuCallback() *FeishuCallback {
	wechatMessageController := impl.NewWechatMessageController()
	return &FeishuCallback{wechatMessageController: wechatMessageController}
}


func (p *FeishuCallback) Callback(c *gin.Context) {
	data := make(map[string]interface{})
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Printf("error: %+v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if msgType, ok := data["type"]; ok {
		if msgType == "url_verification" {
			if receivedToken, ok := data["token"]; ok {
				if receivedToken != configs.DefaultConfig.LarkBot.Token {
					c.JSON(200, gin.H{})
					c.Abort()
					return
				} else {
					if challenge, ok := data["challenge"]; ok {
						c.JSON(http.StatusOK, gin.H{"challenge": challenge})
						c.Abort()
						return
					}
				}
			}
		}
	}
	if action, ok := data["action"]; ok {
		if actionMap, ok := action.(map[string]interface{}); ok {
			log.Printf("%+v\n", actionMap["value"])
			if valueMap, ok := actionMap["value"].(map[string]interface{}); ok {
				wxid, _ := valueMap["wx_id"].(string)
				roomID, _ := valueMap["room_id"].(string)
				createdAt, _ := valueMap["timestamp"].(string)
				direction, _ := valueMap["direction"].(string)
				chatID, _ := valueMap["chat_id"].(string)
				if len(roomID) > 0 && len(createdAt) > 0 && len(direction) > 0 {
					if len(chatID) == 0 {
						p.wechatMessageController.GetRecentMessages(
							configs.DefaultConfig.LarkPictureRoom.ChatID, wxid, roomID, createdAt, direction)
					} else {
						p.wechatMessageController.GetRecentMessages(chatID, wxid, roomID, createdAt, direction)
					}

				}
			}
		}
	}
	c.JSON(http.StatusCreated, gin.H{})
	return
}