package main

import (
	"flag"
	"fmt"
	"github.com/fatelei/juzihudong-sdk/pkg/juzihudong"
	"github.com/fatelei/juzimiaohui-webhook/configs"
	"github.com/fatelei/juzimiaohui-webhook/pkg/connection"
)


func main() {
	var configFile string
	var roomID string
	flag.StringVar(&configFile,"config", "/etc/webhook.toml", "webhook config path")
	flag.StringVar(&roomID,"wxid", "", "room id")
	flag.Parse()

	configs.NewConfig(configFile)
	connection.InitDB()
	roomApi := juzihudong.NewRoomApi(configs.DefaultConfig.Juzihudong.Endpoint, configs.DefaultConfig.Juzihudong.Token)
	current := 0

	for ;; {
		resp := roomApi.GetRooms(current, 10, roomID)
		if len(*resp.Data) == 0 {
			break
		}
		for _, room := range *resp.Data {
			for _, member := range *room.Members {
				fmt.Printf("%s\n", member.Wxid)
			}
		}
		current += 1
	}
}