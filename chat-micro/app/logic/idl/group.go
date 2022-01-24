package idl

import (
	"chat-micro/app/logic/model"
)

//TransferGroupInput 群组对外转化结构
type TransferGroupInput struct {
	Group     *model.GroupModel
	GroupUser []*model.GroupUserModel
	Users     []*model.UserModel
	Nickname  string
}

// TransferGroupInfo 组装数据并输出
// 对外暴露的group结构，都应该经过此结构进行转换
func TransferGroupInfo(input *TransferGroupInput) *model.GroupInfo {
	info := &model.GroupInfo{
		Info: &model.Group{
			ID:            input.Group.ID,
			UserID:        input.Group.UserID,
			InviteConfirm: input.Group.InviteConfirm,
			Name:          input.Group.Name,
			Avatar:        input.Group.Avatar,
			Remark:        input.Group.Remark,
		},
		Nickname: input.Nickname,
		Users:    make([]*model.UserBase, 0),
	}
	uMap := groupUserToMap(input.GroupUser)
	for i, user := range input.Users {
		if i >= 4 { // 最多返回4个用户
			break
		}
		name := user.Username
		if user.Nickname != "" { // 设置了昵称
			name = user.Nickname
		}
		if nick, ok := uMap[user.ID]; ok {
			name = nick
		}
		info.Users = append(info.Users, &model.UserBase{
			ID:     user.ID,
			Name:   name,
			Avatar: user.Avatar,
		})
	}
	return info
}

// TransferGroupUser 组装数据并输出
// 对外暴露的groupUser结构，都应该经过此结构进行转换
func TransferGroupUser(input *TransferGroupInput) []*model.UserBase {
	users := make([]*model.UserBase, 0)
	uMap := groupUserToMap(input.GroupUser)
	for _, user := range input.Users {
		name := user.Username
		if user.Nickname != "" { // 设置了昵称
			name = user.Nickname
		}
		if nick, ok := uMap[user.ID]; ok {
			name = nick
		}
		users = append(users, &model.UserBase{
			ID:     user.ID,
			Name:   name,
			Avatar: user.Avatar,
		})
	}
	return users
}

// 转换群成员map结构 user_id => nickname
func groupUserToMap(gUsers []*model.GroupUserModel) (m map[uint32]string) {
	m = make(map[uint32]string)
	for _, user := range gUsers {
		if user.Nickname != "" {
			m[user.UserID] = user.Nickname
		}
	}
	return
}
