package impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatelei/juzimiaohui-webhook/configs"
	"github.com/fatelei/juzimiaohui-webhook/pkg/controller"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type NotificationControllerImpl struct {
	keywordController *KeywordControllerImpl
}

var _ controller.NotificationController = (*NotificationControllerImpl)(nil)

func NewNotificationController() *NotificationControllerImpl {
	keywordController := NewKeywordController()
	return &NotificationControllerImpl{keywordController:keywordController}
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
			panic(err)
			return
		}
		log.Println(string(rst))
	}
}