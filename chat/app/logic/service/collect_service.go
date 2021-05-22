package service

import (
	"context"

	"github.com/pkg/errors"

	"chat/app/logic/idl"
	"chat/app/logic/model"
	"chat/app/constvar"
)

type ICollect interface {
	// 添加收藏
	CollectCreate(ctx context.Context, content, options string, userId uint32, t int8) error
	// 收藏列表
	CollectGetList(ctx context.Context, userId uint32, offset int) (list []*model.Collect, err error)
	// 删除收藏
	CollectDestroy(ctx context.Context, userId, id uint32) error
}

// CollectCreate 创建收藏
func (s *Service) CollectCreate(ctx context.Context, content, options string, userId uint32, t int8) error {
	collect := &model.CollectModel{
		Uid:     model.Uid{UserId: userId},
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
func (s *Service) CollectGetList(ctx context.Context, userId uint32, offset int) (list []*model.Collect, err error) {
	collectList, err := s.repo.GetCollectsByUserId(ctx, userId, offset, constvar.DefaultLimit)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.apply] GetListByUserId id:%d", userId)
	}
	list = make([]*model.Collect, 0)
	for _, collect := range collectList {
		list = append(list, idl.TransferCollect(collect))
	}
	return list, nil
}

// CollectDestroy 删除收藏
func (s *Service) CollectDestroy(ctx context.Context, userId, id uint32) error {
	err := s.repo.CollectDelete(ctx, userId, id)
	if err != nil {
		return errors.Wrapf(err, "[service.collect] destroy uid:%d,id:%d", userId, id)
	}
	return nil
}
