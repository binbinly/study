package idl

import (
	"chat/app/center/model"
	"chat/proto/base"
	pb "chat/proto/center"
)

//TransferUserInput 用户模型对外转化结构
type TransferUserInput struct {
	User  *model.UserModel
	Token string
}

// TransferAuth 组装数据并输出
// 对外暴露的user auth结构，都应该经过此结构进行转换
func TransferAuth(input *TransferUserInput) *pb.UserToken {
	if input.User == nil {
		return &pb.UserToken{}
	}

	return &pb.UserToken{
		User:  TransferUser(input.User),
		Token: input.Token,
	}
}

// TransferUser 组装数据并输出
// 对外暴露的user结构，都应该经过此结构进行转换
func TransferUser(input *model.UserModel) *base.UserInfo {
	if input == nil {
		return &base.UserInfo{}
	}

	return &base.UserInfo{
		Id:       input.ID,
		Username: input.Username,
		Nickname: input.Nickname,
		Email:    input.Email,
		Avatar:   input.Avatar,
		Sign:     input.Sign,
		Gender:   base.UserInfo_Gender(input.Gender),
		Area:     input.Area,
	}
}
