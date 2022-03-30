package model

import (
	"context"
	"short-link/ent"
	"short-link/ent/user"
)

// CreateUser 创建用户
func CreateUser(u ent.User) (*ent.User, error) {
	return client.User.Create().SetUsername(u.Username).SetRole(u.Role).SetPassword(*u.Password).SetAvatar(u.Avatar).SetPhone(*u.Phone).SetEmail(*u.Email).Save(context.Background())
}

// FindByUsername 根据用户名查询
func FindByUsername(username string) (*ent.User, error) {
	return client.User.Query().Where(user.UsernameEQ(username)).First(context.Background())
}

// UpdateToken 更新用户token
func UpdateToken(id int, t string) (*ent.User, error) {
	return client.User.UpdateOneID(id).SetToken(t).Save(context.Background())
}
