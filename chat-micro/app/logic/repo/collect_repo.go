package repo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"chat-micro/app/logic/model"
	"chat-micro/internal/orm"
	"chat-micro/pkg/logger"
)

//ICollect 收藏接口
type ICollect interface {
	// 创建
	CollectCreate(ctx context.Context, collect *model.CollectModel) (id uint32, err error)
	// 获取用户的收藏列表
	GetCollectsByUserID(ctx context.Context, userID uint32, offset, limit int) (list []*model.CollectModel, err error)
	// 删除收藏
	CollectDelete(ctx context.Context, userID, id uint32) (err error)
}

// CollectCreate 创建收藏
func (r *Repo) CollectCreate(ctx context.Context, collect *model.CollectModel) (id uint32, err error) {
	if err = r.db.WithContext(ctx).Create(collect).Error; err != nil {
		return 0, errors.Wrap(err, "[repo.collect] create collect err")
	}

	r.delCollectCache(ctx, collect.UserID)
	return collect.ID, nil
}

// GetCollectsByUserID 获取用户收藏列表
func (r *Repo) GetCollectsByUserID(ctx context.Context, userID uint32, offset, limit int) (list []*model.CollectModel, err error) {
	if err = r.queryListCache(ctx, collectCacheKey(userID), offset, &list, func(data interface{}) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Scopes(orm.OffsetPage(offset, limit)).Where("user_id = ? ", userID).
			Order(orm.DefaultOrder).Find(data).Error; err != nil {
			return errors.Wrap(err, "[repo.collect] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.collect] query cache")
	}
	return
}

// CollectDelete 删除收藏
func (r *Repo) CollectDelete(ctx context.Context, userID, id uint32) (err error) {
	if err = r.db.WithContext(ctx).Where("user_id=?", userID).Delete(&model.CollectModel{}, id).Error; err != nil {
		return errors.Wrapf(err, "[repo.collect] destroy err")
	}

	r.delCollectCache(ctx, userID)
	return nil
}

//delCollectCache 删除缓存
func (r *Repo) delCollectCache(ctx context.Context, uid uint32) {
	if err := r.cache.Del(ctx, collectCacheKey(uid)); err != nil {
		logger.Warnf("[repo.collect] del cache key: %v", collectCacheKey(uid))
	}
}

func collectCacheKey(uid uint32) string {
	return fmt.Sprintf("collect:list:%d", uid)
}
