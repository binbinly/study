package idl

import "chat/app/logic/model"

type TransferApplyInput struct {
	Apply []*model.ApplyModel
	Users []*model.UserModel
}

// TransferApplyList 组装数据并输出
// 对外暴露的apply结构，都应该经过此结构进行转换
func TransferApplyList(input *TransferApplyInput) []*model.ApplyList {
	um := usersToMap(input.Users)
	list := make([]*model.ApplyList, 0)
	for _, apply := range input.Apply {
		if user, ok := um[apply.UserId]; ok {
			list = append(list, &model.ApplyList{
				User:   user,
				Status: apply.Status,
			})
		}
	}
	return list
}