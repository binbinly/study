package idl

import (
	"chat/app/chat/model"
	"chat/proto/base"
)

// TransferUserBase 组装数据并输出
// 对外暴露的user结构，都应该经过此结构进行转换
func TransferUserBase(input *base.UserInfo) *model.UserBase {
	if input == nil {
		return &model.UserBase{}
	}

	name := input.Username
	if input.Nickname != "" {
		name = input.Nickname
	}
	return &model.UserBase{
		ID:     input.Id,
		Name:   name,
		Avatar: input.Avatar,
	}
}

//TransferUserName 获取用户显示name
func TransferUserName(input *base.UserInfo) string {
	name := input.Username
	if input.Nickname != "" {
		name = input.Nickname
	}
	return name
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
