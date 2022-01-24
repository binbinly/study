package service

import (
	"context"

	"github.com/pkg/errors"

	"chat-micro/app/constvar"
	"chat-micro/app/logic/idl"
	"chat-micro/app/logic/model"
	"chat-micro/internal/orm"
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
		UID:     orm.UID{UserID: userID},
		Content: content,
		Type:    t,
		Options: options,
	}
	if _, err := s.repo.CollectCreate(ctx, collect); err != nil {
		return errors.Wrapf(err, "[service.collect] create err")
	}
	return nil
}

// CollectGetList 我的收藏列表
func (s *Service) CollectGetList(ctx context.Context, userID uint32, offset int) ([]*model.Collect, error) {
	list, err := s.repo.GetCollectsByUserID(ctx, userID, offset, constvar.DefaultLimit)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.collect] id:%d", userID)
	}
	return idl.TransferCollectList(list), nil
}

// CollectDestroy 删除收藏
func (s *Service) CollectDestroy(ctx context.Context, userID, id uint32) error {
	err := s.repo.CollectDelete(ctx, userID, id)
	if err != nil {
		return errors.Wrapf(err, "[service.collect] destroy uid:%d,id:%d", userID, id)
	}
	return nil
}
