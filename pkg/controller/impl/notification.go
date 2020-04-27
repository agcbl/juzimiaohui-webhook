package impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatelei/go-feishu/pkg/image"
	"github.com/fatelei/go-feishu/pkg/message"
	feishuModel "github.com/fatelei/go-feishu/pkg/model/interactive"
	"github.com/fatelei/juzimiaohui-webhook/configs"
	"github.com/fatelei/juzimiaohui-webhook/pkg/controller"
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao"
	"github.com/fatelei/juzimiaohui-webhook/pkg/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type NotificationControllerImpl struct {
	keywordController   controller.KeywordController
	feishuBotController controller.FeishuBotController
	feishuMessageApi    *message.MessageAPI
	feishuImageApi      *image.ImageAPI
}

var _ controller.NotificationController = (*NotificationControllerImpl)(nil)

func NewNotificationController() *NotificationControllerImpl {
	keywordController := NewKeywordController()
	feishuBotController := NewFeishuBotController()
	feishuMessageApi := message.NewMessageAPI(configs.DefaultConfig.LarkBot.EndPoint)
	feishuImageApi := image.NewImageAPI(configs.DefaultConfig.LarkBot.EndPoint)
	return &NotificationControllerImpl{
		keywordController: keywordController, feishuMessageApi: feishuMessageApi, feishuImageApi: feishuImageApi, feishuBotController: feishuBotController}
}

func (p *NotificationControllerImpl) CreateMessageCard(message *model.WechatMessage) {
	accessToken := p.feishuBotController.GetAccessToken()
	imageResp, err := p.feishuImageApi.UploadFromUri(message.Payload.ImageUrl, accessToken)
	if err == nil && imageResp.Data != nil {
		prevButton := feishuModel.ButtonModule{
			Tag:   "button",
			Text:  &feishuModel.TextModule{Tag: "plain_text", Content: "获取用户前 10 条消息"},
			Value: make(map[string]string),
		}
		prevButton.SetValue("wx_id", message.ContactId)
		prevButton.SetValue("room_id", message.RoomId)
		prevButton.SetValue("timestamp", strconv.Itoa(message.Timestamp))
		prevButton.SetValue("direction", "before")
		prevButton.SetValue("chat_id", configs.DefaultConfig.LarkPictureRoom.ChatID)
		prevButton.SetValue("action", "loadMyRoomMessage")

		nextButton := feishuModel.ButtonModule{
			Tag:   "button",
			Text:  &feishuModel.TextModule{Tag: "plain_text", Content: "获取用户后 10 条消息"},
			Value: make(map[string]string),
		}
		nextButton.SetValue("wx_id", message.ContactId)
		nextButton.SetValue("room_id", message.RoomId)
		nextButton.SetValue("timestamp", strconv.Itoa(message.Timestamp))
		nextButton.SetValue("direction", "after")
		nextButton.SetValue("chat_id", configs.DefaultConfig.LarkPictureRoom.ChatID)
		nextButton.SetValue("action", "loadMyRoomMessage")

		actionModule := &feishuModel.ActionModule{
			Tag:     "action",
			Actions: []feishuModel.Interactive{prevButton, nextButton},
		}
		title := fmt.Sprintf("%s-%s-%s-%s-%s", message.ContactId, message.ContactName, message.RoomTopic, message.RoomId, time.Now().Format("2006-01-02 15:04:05"))
		p.feishuMessageApi.SendImage(
			configs.DefaultConfig.LarkPictureRoom.ChatID, title, imageResp.Data.ImageKey, actionModule, accessToken)
	}
}

func (p *NotificationControllerImpl) SendRecentMessagesCard(chatID string, messages []*dao.WechatMessage) {
	accessToken := p.feishuBotController.GetAccessToken()
	if len(messages) > 0 {
		elements := make([]interface{}, len(messages))
		var title string
		for index, item := range messages {
			if len(title) == 0 {
				title = fmt.Sprintf("%s（%s）在群「%s」中说：", item.WxID, item.WechatName, item.RoomName)
			}
			element := &feishuModel.ContentModule{
				Tag: "div",
				Text: &feishuModel.TextModule{
					Tag:     "plain_text",
					Content: fmt.Sprintf("%s %s", item.CreatedAt.Format("2006-01-02 15:04:05"), item.Content),
				},
			}
			elements[index] = element
		}
		resp, err := p.feishuMessageApi.SendInteractiveCard(chatID, title, elements, accessToken)
		if err != nil {
			log.Printf("%+v\n", err)
		} else {
			log.Printf("%+v\n", resp)
		}
	}
}


func (p *NotificationControllerImpl) SendRoomRecentMessagesCard(
	chatID string, messages []*dao.WechatMessage, roomAlias map[string]string) {
	accessToken := p.feishuBotController.GetAccessToken()
	if len(messages) > 0 {
		elements := make([]interface{}, len(messages))
		var title string
		var hasAlias bool
		var alias string
		for index, item := range messages {
			if len(title) == 0 {
				title = fmt.Sprintf("%s - %s：", item.RoomName, item.RoomID)
			}

			alias, hasAlias = roomAlias[item.WxID]
			if !hasAlias {
				alias = item.WechatName
			}

			element := &feishuModel.ContentModule{
				Tag: "div",
				Text: &feishuModel.TextModule{
					Tag:     "plain_text",
					Content: fmt.Sprintf(
						"%s %s(%s)：%s", item.CreatedAt.Format("2006-01-02 15:04:05"), alias, item.WxID, item.Content),
				},
			}
			elements[index] = element
		}
		resp, err := p.feishuMessageApi.SendInteractiveCard(chatID, title, elements, accessToken)
		if err != nil {
			log.Printf("%+v\n", err)
		} else {
			log.Printf("%+v\n", resp)
		}
	}
}


func (p *NotificationControllerImpl) CreateNotification(wechatMessage *model.WechatMessage) {
	content := wechatMessage.GetContent()
	hitWord := p.keywordController.Search(content)
	fmt.Printf("hit words: %s\n", hitWord)
	if len(hitWord) > 0 && configs.DefaultConfig.Lark != nil {
		index := strings.Index(content, hitWord)
		var sendContent string

		prevButton := feishuModel.ButtonModule{
			Tag:   "button",
			Text:  &feishuModel.TextModule{Tag: "plain_text", Content: "获取该群前 10 条消息"},
			Value: make(map[string]string),
		}
		prevButton.SetValue("wx_id", "")
		prevButton.SetValue("room_id", wechatMessage.RoomId)
		prevButton.SetValue("timestamp", strconv.Itoa(wechatMessage.Timestamp))
		prevButton.SetValue("direction", "before")
		prevButton.SetValue("chat_id", configs.DefaultConfig.LarkTextRoom.ChatID)
		prevButton.SetValue("action", "loadRoomMessage")

		nextButton := feishuModel.ButtonModule{
			Tag:   "button",
			Text:  &feishuModel.TextModule{Tag: "plain_text", Content: "获取该群后 10 条消息"},
			Value: make(map[string]string),
		}
		nextButton.SetValue("wx_id", "")
		nextButton.SetValue("room_id", wechatMessage.RoomId)
		nextButton.SetValue("timestamp", strconv.Itoa(wechatMessage.Timestamp))
		nextButton.SetValue("direction", "after")
		nextButton.SetValue("chat_id", configs.DefaultConfig.LarkTextRoom.ChatID)
		nextButton.SetValue("action", "loadRoomMessage")

		actionModule := &feishuModel.ActionModule{
			Tag:     "action",
			Actions: []feishuModel.Interactive{prevButton, nextButton},
		}
		if index != -1 {
			sendContent = fmt.Sprintf("%s『%s』%s", content[:index], hitWord, content[index+len(hitWord):])
		} else {
			sendContent = content
		}

		elements := make([]interface{}, 2)
		elements[0] = &feishuModel.ContentModule{
			Tag: "div",
			Text: &feishuModel.TextModule{
				Tag:     "plain_text",
				Content: sendContent,
			},
		}
		elements[1] = actionModule

		title := fmt.Sprintf(
			"%s（%s）在群「%s」中说：", wechatMessage.ContactName, wechatMessage.ContactId, wechatMessage.RoomTopic)
		accessToken := p.feishuBotController.GetAccessToken()
		p.feishuMessageApi.SendInteractiveCard(
			configs.DefaultConfig.LarkTextRoom.ChatID, title, elements, accessToken)
	}
}

func (p *NotificationControllerImpl) CreateWechatDeathNoti() {
	var body []byte
	var err error
	var rst []byte
	var resp *http.Response

	if configs.DefaultConfig.Lark != nil {
		body, err = json.Marshal(map[string]string{
			"title": "最近 30 分钟内无消息",
			"text":  "请尽快检查绑定微信是否下线，如果已下线请尽快前往<a href=\"https://wechat.botorange.com/\">句子互动</a>后台重新登录",
		})
		if err != nil {
			panic(err)
			return
		}
		resp, err = http.Post(
			configs.DefaultConfig.Lark.Path, "application/json", bytes.NewReader(body))
		if err != nil {
			panic(err)
			return
		}
		defer resp.Body.Close()
		rst, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("error %+v\n", err)
			return
		}
		log.Println(string(rst))
	}
}
