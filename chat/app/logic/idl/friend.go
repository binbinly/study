package idl

import "chat/app/logic/model"

type TransferFriendInput struct {
	User       *model.UserModel
	Friend     *model.FriendModel
	FriendTags []string
}

// TransferFriend 组装数据并输出
// 对外暴露的friend结构，都应该经过此结构进行转换
func TransferFriend(input *TransferFriendInput) *model.FriendInfo {
	f := &model.FriendInfo{
		User: TransferUser(input.User),
	}
	if input.Friend.ID == 0 {
		f.Friend = nil
	} else {
		f.IsFriend = true
		f.Friend = &model.FriendBase{
			Nickname: input.Friend.Nickname,
			LookMe:   input.Friend.LookMe,
			LookHim:  input.Friend.LookHim,
			IsStar:   input.Friend.IsStar,
			IsBlack:  input.Friend.IsBlack,
			Tags:     input.FriendTags,
		}
	}
	return f
}
