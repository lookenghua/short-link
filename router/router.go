package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	v1 "short-link/router/api/v1"
)

// InitRouter 初始化路由
func InitRouter() *fiber.App {
	app := fiber.New()

	// 全局中间件
	{
		// 异常中间件
		app.Use(recover.New())
		// 压缩中间件
		app.Use(compress.New(compress.Config{
			Level: compress.LevelBestSpeed, // 1
		}))
		// 跨域中间件
		app.Use(cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
		}))
		// 日志中间件
		app.Use(logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		}))
		// RequestID中间件
		app.Use(requestid.New(requestid.Config{
			Header: "X-Custom-Header",
			Generator: func() string {
				return "static-id"
			},
		}))

	}

	// 路由
	{
		app.Get("/", func(ctx *fiber.Ctx) error {
			return ctx.SendString("start success")
		})
		api := app.Group("/api")
		apiV1 := api.Group("/v1")
		apiV1.Post("/user", v1.CreateUser)
	}

	return app
}
