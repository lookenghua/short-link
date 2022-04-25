package guard

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"short-link/service"
	"short-link/util"
	"strings"
)

func UserGuard(c *fiber.Ctx) error {
	var headers = c.GetReqHeaders()
	var bearerToken = headers["Authorization"]
	apiUtil := util.ApiUtil{Ctx: c}
	if bearerToken == "" {
		return apiUtil.Fail(fiber.StatusUnauthorized, "请先登录")
	}
	if !strings.HasPrefix(bearerToken, "bearer") {
		return apiUtil.Fail(util.IllegalToken, "token不合法")
	}
	var token = bearerToken[7:]
	user, findErr := service.FindByToken(token)
	if findErr != nil {
		return apiUtil.Fail(util.DataFoundError, "token不合法")
	}
	if user == nil {
		return apiUtil.Fail(util.DataNotFoundError, "用户不存在")
	}
	fmt.Printf("当前访问用户ID: %v \n", user.ID)
	//// 存储用户数据到全局
	apiUtil.SetUserInfo(user)
	return c.Next()
}
