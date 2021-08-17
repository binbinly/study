package idl

import (
	"chat/app/chat/model"
	"chat/proto/base"
)

// 用户列表转换成map结构
func usersToMap(users []*base.UserInfo) (m map[uint32]*model.UserBase) {
	m = make(map[uint32]*model.UserBase)
	for _, user := range users {
		name := user.Username
		if user.Nickname != "" {
			name = user.Nickname
		}
		m[user.Id] = &model.UserBase{
			ID:     user.Id,
			Name:   name,
			Avatar: user.Avatar,
		}
	}
	return m
}
