package idl

import (
	"center/model"
	pb "common/proto/center"
)

// TransferUser 组装数据并输出
// 对外暴露的user结构，都应该经过此结构进行转换
func TransferUser(input *model.UserModel) *pb.Userinfo {
	if input == nil {
		return &pb.Userinfo{}
	}

	return &pb.Userinfo{
		Id:       input.ID,
		Username: input.Username,
		Nickname: input.Nickname,
		Phone:    input.Phone,
		Email:    input.Email,
		Avatar:   input.Avatar,
		Sign:     input.Sign,
		Gender:   pb.Userinfo_Gender(input.Gender),
		Area:     input.Area,
	}
}
