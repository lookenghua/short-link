package v1

import (
	"github.com/gofiber/fiber/v2"
	. "short-link/common/dto"
	"short-link/service"
	. "short-link/util"
)

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

// LoginUser 用户登录
func LoginUser(c *fiber.Ctx) error {
	apiUtil := ApiUtil{Ctx: c}
	data := new(LoginUserDto)
	err := apiUtil.ValidateBody(data)
	if err != nil {
		return apiUtil.Fail(ValidateError, err.Error())
	}
	// 判断用户存在和密码正确性
	user, findErr := service.FindByUsername(data.Username)
	if findErr != nil {
		return apiUtil.Fail(QUERYError, findErr.Error())
	}
	if user == nil {
		return apiUtil.Fail(DataFoundError, "用户名不存在")
	}
	// TODO:MD5加密
	if *user.Password != data.Password {
		return apiUtil.Fail(PasswordNotMatchedError, "密码错误")
	}
	// 创建token
	token, createErr := service.CreateToken(user.ID)
	if createErr != nil {
		return apiUtil.Fail(CreateTokenError, findErr.Error())
	}
	// 存入token
	_, saveErr := service.SaveToken(user.ID, token)
	if saveErr != nil {
		return apiUtil.Fail(SaveDataError, findErr.Error())
	}

	return apiUtil.Success(token)
}
