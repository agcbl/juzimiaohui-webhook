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

type NotificationControllerImpl struct {}

var DefaultNotificationController *NotificationControllerImpl

var _ controller.NotificationController = (*NotificationControllerImpl)(nil)


func (p *NotificationControllerImpl) CreateNotification(room string, contact string, content string) {
	var body []byte
	var err error
	var rst []byte
	var resp *http.Response
	var flag = false
	for _, word := range configs.DefaultConfig.Word.Words {
		if strings.Contains(content, word) {
			flag = true
			break
		}
	}

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