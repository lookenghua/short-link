package service

import (
	"errors"
	. "short-link/common/dto"
	"short-link/ent"
	"short-link/ent/user"
	"short-link/model"
)
import . "short-link/util"

// CreateUser 创建用户
func CreateUser(data *CreateUserDto) (*ent.User, error) {
	u, err := model.FindByUsername(data.Username)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if u != nil {
		return nil, errors.New("用户已存在")
	}
	// 用户不存在
	userInfo := ent.User{
		Username: data.Username,
		Password: &data.Password,
		Role:     user.RoleUSER,
		Avatar:   data.Avatar,
		Phone:    &data.Phone,
		Email:    &data.Email,
	}
	newUser, createErr := model.CreateUser(userInfo)
	if createErr != nil {
		return nil, createErr
	}
	return newUser, nil
}
