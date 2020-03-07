package main

import (
	"flag"
	"fmt"
	"github.com/fatelei/juzihudong-sdk/pkg/juzihudong"
	"github.com/fatelei/juzimiaohui-webhook/configs"
	"github.com/fatelei/juzimiaohui-webhook/pkg/connection"
	"github.com/fatelei/juzimiaohui-webhook/pkg/controller/impl"
	"log"
)


func main() {
	var endpoint string
	var token string
	var configFile string
	flag.StringVar(&configFile,"config", "/etc/webhook.toml", "webhook config path")
	flag.StringVar(&endpoint, "endpoint", "https://ex-api.botorange.com", "api endpoint")
	flag.StringVar(&token, "token", "","api token")
	flag.Parse()

	configs.NewConfig(configFile)
	connection.InitDB()
	if len(token) == 0 {
		fmt.Println("token is required")
		return
	}
	roomApi := juzihudong.NewRoomApi(endpoint, token)
	current := 106

	for ;; {
		log.Printf("sync room, page = %d\n", current)
		resp := roomApi.GetRooms(current, 10)
		if len(*resp.Data) == 0 {
			break
		} else {
			for _, room := range *resp.Data {
				impl.DefaultWechatRoomController.CreateRoom(&room)
			}
		}
		current += 1
	}
}