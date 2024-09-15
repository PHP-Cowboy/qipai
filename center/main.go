package main

import (
	"center/app"
	"common/config"
	"common/initialize"
	"context"
	"frame/game"
	"log"
	"os"
)

func main() {
	initialize.InitLogger()

	config.InitConfig("config.json")

	game.InitConfig("../config")

	serverId := "001"

	//启动服务端
	err := app.Run(context.Background(), serverId)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}
