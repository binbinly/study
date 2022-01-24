package idl

import "chat-micro/app/logic/model"

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

//TransferCollectList 转化列表输出
func TransferCollectList(list []*model.CollectModel) []*model.Collect {
	if len(list) == 0 {
		return []*model.Collect{}
	}
	res := make([]*model.Collect, 0, len(list))
	for _, collect := range list {
		res = append(res, TransferCollect(collect))
	}
	return res
}