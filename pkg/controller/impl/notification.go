package impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatelei/go-feishu/pkg/message"
	"github.com/fatelei/go-feishu/pkg/image"
	feishuModel "github.com/fatelei/go-feishu/pkg/model/interactive"
	"github.com/fatelei/juzimiaohui-webhook/configs"
	"github.com/fatelei/juzimiaohui-webhook/pkg/controller"
	"github.com/fatelei/juzimiaohui-webhook/pkg/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type NotificationControllerImpl struct {
	keywordController *KeywordControllerImpl
	feishuMessageApi *message.MessageAPI
	feishuImageApi *image.ImageAPI
}

var _ controller.NotificationController = (*NotificationControllerImpl)(nil)

func NewNotificationController() *NotificationControllerImpl {
	keywordController := NewKeywordController()
	feishuMessageApi := message.NewMessageAPI(
		configs.DefaultConfig.LarkBot.AppID, configs.DefaultConfig.LarkBot.AppSecret, configs.DefaultConfig.LarkBot.EndPoint)
	feishuImageApi := image.NewImageAPI(
		configs.DefaultConfig.LarkBot.AppID, configs.DefaultConfig.LarkBot.AppSecret, configs.DefaultConfig.LarkBot.EndPoint)
	return &NotificationControllerImpl{
		keywordController:keywordController, feishuMessageApi: feishuMessageApi, feishuImageApi: feishuImageApi}
}

func (p *NotificationControllerImpl) CreateMessageCard(message model.WechatMessage) {
	imageResp, err := p.feishuImageApi.UploadFromUri(message.Payload.ImageUrl)
	if err == nil && imageResp.Data != nil {
		prevButton := feishuModel.ButtonModule{
			Tag:   "button",
			Text:  &feishuModel.TextModule{Tag: "plain_text", Content: "获取用户前 10 条消息"},
			Value: make(map[string]string),
		}
		prevButton.SetValue("wxid", message.ContactId)
		prevButton.SetValue("timestamp", strconv.Itoa(message.Timestamp))

		nextButton := feishuModel.ButtonModule{
			Tag:   "button",
			Text:  &feishuModel.TextModule{Tag: "plain_text", Content: "获取用户后 10 条消息"},
			Value: make(map[string]string),
		}
		nextButton.SetValue("wxid", message.ContactId)
		nextButton.SetValue("timestamp", strconv.Itoa(message.Timestamp))

		actionModule := &feishuModel.ActionModule{
			Tag:     "action",
			Actions: []feishuModel.Interactive{prevButton, nextButton},
		}
		title := fmt.Sprintf("%s-%s-%s-%s-%s", message.ContactId, message.ContactName, message.RoomTopic, message.RoomId, time.Now().Format("2006-01-02 15:04:05"))
		p.feishuMessageApi.SendImage(configs.DefaultConfig.LarkBot.ChatID, title, imageResp.Data.ImageKey, actionModule)
	}
}


func (p *NotificationControllerImpl) CreateNotification(room string, contactName string, contactId string, content string) {
	var body []byte
	var err error
	var rst []byte
	var resp *http.Response
	hitWord := p.keywordController.Search(content)
	fmt.Printf("hit words: %s\n", hitWord)
	if len(hitWord) > 0 && configs.DefaultConfig.Lark != nil {
		index := strings.Index(content, hitWord)
		var sendContent string
		if index != -1 {
			sendContent = fmt.Sprintf("%s『%s』%s", content[:index], hitWord, content[index + len(hitWord):])
		} else {
			sendContent = content
		}
		body, err = json.Marshal(map[string]string{
			"title": fmt.Sprintf("%s（%s）在群「%s」中说：", contactName, contactId, room),
			"text": fmt.Sprintf("%s", sendContent),
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
			panic(err)
			return
		}
		log.Println(string(rst))
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
			"text": "请尽快检查绑定微信是否下线，如果已下线请尽快前往<a href=\"https://wechat.botorange.com/\">句子互动</a>后台重新登录",
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