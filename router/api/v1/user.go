package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"short-link/ent"
	"short-link/ent/user"
	"short-link/model"
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

	u, err := model.FindByUsername(data.Username)
	if err != nil && !ent.IsNotFound(err) {
		return apiUtil.Fail(UnknownError, err.Error())
	} else if u != nil {
		return apiUtil.Fail(DataFoundError, "用户名已存在")
	}
	// 用户不存在
	user := ent.User{
		Username: data.Username,
		Password: &data.Password,
		Role:     user.RoleUSER,
		Avatar:   data.Avatar,
		Phone:    &data.Phone,
		Email:    &data.Email,
	}
	newUser, createErr := model.CreateUser(user)
	if createErr != nil {
		return apiUtil.Fail(CreateError, createErr.Error())
	}
	fmt.Printf("%v\n", user)

	return apiUtil.Success(newUser)
}
