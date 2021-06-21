package service

import (
	"context"

	"github.com/pkg/errors"

	"chat/app/constvar"
	"chat/app/logic/idl"
	"chat/app/logic/model"
)

//ICollect 用户收藏服务接口
type ICollect interface {
	// 添加收藏
	CollectCreate(ctx context.Context, content, options string, userID uint32, t int8) error
	// 收藏列表
	CollectGetList(ctx context.Context, userID uint32, offset int) (list []*model.Collect, err error)
	// 删除收藏
	CollectDestroy(ctx context.Context, userID, id uint32) error
}

// CollectCreate 创建收藏
func (s *Service) CollectCreate(ctx context.Context, content, options string, userID uint32, t int8) error {
	collect := &model.CollectModel{
		UID:     model.UID{UserID: userID},
		Content: content,
		Type:    t,
		Options: options,
	}
	_, err := s.repo.CollectCreate(ctx, collect)
	if err != nil {
		return errors.Wrapf(err, "[service.collect] create err")
	}
	return nil
}

// CollectGetList 我的收藏列表
func (s *Service) CollectGetList(ctx context.Context, userID uint32, offset int) (list []*model.Collect, err error) {
	collectList, err := s.repo.GetCollectsByUserID(ctx, userID, offset, constvar.DefaultLimit)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.apply] GetListByUserId id:%d", userID)
	}
	list = make([]*model.Collect, 0)
	for _, collect := range collectList {
		list = append(list, idl.TransferCollect(collect))
	}
	return list, nil
}

// CollectDestroy 删除收藏
func (s *Service) CollectDestroy(ctx context.Context, userID, id uint32) error {
	err := s.repo.CollectDelete(ctx, userID, id)
	if err != nil {
		return errors.Wrapf(err, "[service.collect] destroy uid:%d,id:%d", userID, id)
	}
	return nil
}
