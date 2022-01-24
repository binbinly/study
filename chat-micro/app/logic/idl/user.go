package idl

import (
	"chat-micro/app/logic/model"
)

// TransferUser 组装数据并输出
// 对外暴露的user结构，都应该经过此结构进行转换
func TransferUser(input *model.UserModel) *model.User {
	if input == nil {
		return &model.User{}
	}

	return &model.User{
		ID:       input.ID,
		Phone:    input.Phone,
		Username: input.Username,
		Nickname: input.Nickname,
		Email:    input.Email,
		Avatar:   input.Avatar,
		Sign:     input.Sign,
		Area:     input.Area,
		Gender:   input.Gender,
	}
}

//TransferUsers 输出用户列表
func TransferUsers(input []*model.UserModel) []*model.User {
	if len(input) == 0 {
		return []*model.User{}
	}

	users := make([]*model.User, 0, len(input))
	for _, user := range input {
		users = append(users, TransferUser(user))
	}
	return users
}

// TransferUserBase 组装数据并输出
// 对外暴露的user结构，都应该经过此结构进行转换
func TransferUserBase(input *model.UserModel) *model.UserBase {
	if input == nil {
		return &model.UserBase{}
	}

	return &model.UserBase{
		ID:     input.ID,
		Name:   TransferUserName(input),
		Avatar: input.Avatar,
	}
}

//TransferUserName 获取用户显示name
func TransferUserName(input *model.UserModel) string {
	name := input.Username
	if input.Nickname != "" {
		name = input.Nickname
	}
	return name
}
