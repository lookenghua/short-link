package router

import (
	"github.com/goccy/go-json"
	"short-link/router/guard"
	. "short-link/router/middleware"
)
import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	v1 "short-link/router/api/v1"
)

// InitRouter 初始化路由
func InitRouter() *fiber.App {
	// 自定义json转换器,提高效率
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

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
		// 日志记录中间件
		app.Use(LoggerMiddleware)
	}
	// 全局路由
	{
		// 展示监控
		app.Get("/", monitor.New())
	}

	// 路由
	{
		api := app.Group("/api")
		apiV1 := api.Group("/v1")
		// 创建用户
		apiV1.Post("/user", v1.CreateUser)
		// 用户登录
		apiV1.Post("/user/token", v1.LoginUser)
		// 获取用户登录信息
		apiV1.Get("/user", guard.UserGuard, v1.GetCurrentUserInfo)
	}

	return app
}
