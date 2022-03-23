package main

import (
	"log"
	"short-link/common/validation"
	"short-link/model"
	"short-link/router"
)

func init() {
	model.InitDatabase()
	validation.InitValidation()
}

func main() {
	app := router.InitRouter()
	err := app.Listen(":6605")
	if err != nil {
		log.Println("服务器启动失败")
		return
	}
	log.Println("服务器启动成功")
}
