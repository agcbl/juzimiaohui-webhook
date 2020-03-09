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
)

type NotificationControllerImpl struct {
	keywordController *KeywordControllerImpl
}

var _ controller.NotificationController = (*NotificationControllerImpl)(nil)

func NewNotificationController() *NotificationControllerImpl {
	keywordController := NewKeywordController()
	return &NotificationControllerImpl{keywordController:keywordController}
}


func (p *NotificationControllerImpl) CreateNotification(room string, contact string, content string) {
	var body []byte
	var err error
	var rst []byte
	var resp *http.Response
	flag := p.keywordController.Search(content)

	if flag {
		body, err = json.Marshal(map[string]string{
			"title": fmt.Sprintf("微信群 %s 反馈", room),
			"text": fmt.Sprintf("%s: %s", contact, content),
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