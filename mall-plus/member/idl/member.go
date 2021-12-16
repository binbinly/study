package idl

import (
	"strconv"

	center "common/proto/center"
	member "common/proto/member"
	"member/model"
)

//TransferMemberToken 转化会员登录成功信息
func TransferMemberToken(user *center.Userinfo) *member.MemberInfo {
	return &member.MemberInfo{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Sign:     user.Sign,
		Avatar:   user.Avatar,
		Area:     user.Area,
		Phone:    strconv.FormatInt(user.Phone, 10),
	}
}

// TransferMember 会员信息转换输出
func TransferMember(model *model.MemberModel) *member.MemberInfo {
	return &member.MemberInfo{
		Id:       model.ID,
		Username: model.Username,
		Nickname: model.Nickname,
		Sign:     model.Sign,
		Avatar:   model.Avatar,
		Area:     model.Area,
		Phone:    strconv.FormatInt(model.Phone, 10),
	}
}
