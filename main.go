package main

import (
	"log"
	"short-link/common/validation"
	"short-link/model"
	"short-link/router"
	"short-link/util"
)

func init() {
	util.InitLogger()
	model.InitDatabase()
	validation.InitValidation()
}

func main() {
	app := router.InitRouter()
	err := app.Listen(":6605")
	if err != nil {
		util.Logger.Error("服务器启动失败")
		return
	}
	log.Println("服务器启动成功")
}
