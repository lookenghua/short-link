package main

import (
	"log"
	"short-link/router"
)

func main() {
	app := router.InitRouter()
	err := app.Listen(":6605")
	if err != nil {
		log.Println("服务器启动失败")
		return
	}
	log.Println("服务器启动成功")
}
