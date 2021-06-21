package idl

import "chat/app/logic/model"

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
		User: TransferUser(input.User),
		Token: input.Token,
	}
}

// TransferUserBase 组装数据并输出
// 对外暴露的user结构，都应该经过此结构进行转换
func TransferUserBase(input *model.UserModel) *model.UserBase {
	if input == nil {
		return &model.UserBase{}
	}

	name := input.Username
	if input.Nickname != "" {
		name = input.Nickname
	}
	return &model.UserBase{
		ID:     input.ID,
		Name:   name,
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

// TransferUser 组装数据并输出
// 对外暴露的user结构，都应该经过此结构进行转换
func TransferUser(input *model.UserModel) *model.UserInfo {
	if input == nil {
		return &model.UserInfo{}
	}

	return &model.UserInfo{
		ID:       input.ID,
		Username: input.Username,
		Nickname: input.Nickname,
		Avatar:   input.Avatar,
		Sign:     input.Sign,
		Gender:   input.Gender,
	}
}

// TransferCollect 组装数据并输出
// 对外暴露的collect结构，都应该经过此结构进行转换
func TransferCollect(input *model.CollectModel) *model.Collect {
	if input.ID == 0 {
		return &model.Collect{}
	}

	return &model.Collect{
		ID:      input.ID,
		Type:    input.Type,
		Content: input.Content,
		Options: input.Options,
	}
}