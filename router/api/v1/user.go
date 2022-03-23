package v1

import (
	"github.com/gofiber/fiber/v2"
	"short-link/service"
)
import . "short-link/common/dto"
import . "short-link/util"

// CreateUser 创建用户
func CreateUser(c *fiber.Ctx) error {
	apiUtil := ApiUtil{Ctx: c}
	data := new(CreateUserDto)
	err := apiUtil.ValidateBody(data)
	if err != nil {
		return apiUtil.Fail(ValidateError, err.Error())
	}
	newUser, createErr := service.CreateUser(data)
	if createErr != nil {
		return apiUtil.Fail(CreateError, err.Error())
	}
	return apiUtil.Success(newUser)
}
