package middleware

import (
	"github.com/gofiber/fiber/v2"
	"short-link/util"
	"time"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	startTime := time.Now()
	err := c.Next()
	endTime := time.Now()
	t := endTime.Sub(startTime)
	if t.Seconds() > 5 {
		// 请求大于5秒记录到日志中
		util.Logger.Error("请求时间过长", " ", c.Method(), " ", c.Path())
	}

	return err
}
