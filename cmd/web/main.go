package main

import (
	"flag"
	"fmt"
	"github.com/fatelei/juzimiaohui-webhook/configs"
	"github.com/fatelei/juzimiaohui-webhook/pkg/connection"
	"github.com/fatelei/juzimiaohui-webhook/web"
	"log"
	"time"
)

func main() {
	var port int
	var configFile string
	log.SetFlags(log.Ldate)
	log.SetFlags(log.Ltime)
	flag.IntVar(&port, "port", 8000, "http server port")
	flag.StringVar(&configFile,"config", "/etc/webhook.toml", "webhook config path")
	flag.Parse()

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	time.Local = loc
	configs.NewConfig(configFile)
	connection.InitDB()
	engine := web.Routes()
	err = engine.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
