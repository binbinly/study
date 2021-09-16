package idl

import "mall/app/model"

//TransferUserInput 用户模型对外转化结构
type TransferUserInput struct {
	User  *model.UserModel
	Token string
}

// TransferAuth 组装数据并输出
// 对外暴露的user auth结构，都应该经过此结构进行转换
func TransferAuth(input *TransferUserInput) *model.UserToken {
	if input.User == nil {
		return &model.UserToken{}
	}

	return &model.UserToken{
		User:  TransferUser(input.User),
		Token: input.Token,
	}
}

// TransferUser 组装数据并输出
// 对外暴露的user结构，都应该经过此结构进行转换
func TransferUser(input *model.UserModel) *model.Userinfo {
	if input == nil {
		return &model.Userinfo{}
	}

	return &model.Userinfo{
		ID:       input.ID,
		Username: input.Username,
		Nickname: input.Nickname,
		Avatar:   input.Avatar,
	}
}
