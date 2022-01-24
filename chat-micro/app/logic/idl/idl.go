package idl

import (
	"chat-micro/app/logic/model"
)

// 用户列表转换成map结构
func usersToMap(users []*model.UserModel) (m map[uint32]*model.UserBase) {
	m = make(map[uint32]*model.UserBase)
	for _, user := range users {
		name := user.Username
		if user.Nickname != "" {
			name = user.Nickname
		}
		m[user.ID] = &model.UserBase{
			ID:     user.ID,
			Name:   name,
			Avatar: user.Avatar,
		}
	}
	return m
}
