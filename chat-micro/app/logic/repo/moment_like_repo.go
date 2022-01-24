package repo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"chat-micro/app/logic/model"
	"chat-micro/pkg/logger"
	"chat-micro/pkg/util"
)

//IMomentLike 朋友圈点赞
type IMomentLike interface {
	// 创建
	LikeCreate(ctx context.Context, model *model.MomentLikeModel) (id uint32, err error)
	// 删除数据
	LikeDelete(ctx context.Context, user, moment uint32) error
	// 是否存在
	LikeExist(ctx context.Context, userID, momentID uint32) (bool, error)
	// 朋友圈动态点赞用户列表
	GetLikeUserIdsByMomentID(ctx context.Context, momentID uint32) ([]uint32, error)
	// 朋友圈动态点赞列表
	GetLikesByMomentIds(ctx context.Context, momentIds []uint32) (map[uint32][]uint32, error)
}

// LikeCreate 创建
func (r *Repo) LikeCreate(ctx context.Context, model *model.MomentLikeModel) (id uint32, err error) {
	if err = r.db.WithContext(ctx).Create(model).Error; err != nil {
		return 0, errors.Wrapf(err, "[repo.moment_like] create err")
	}
	r.delLikeCache(ctx, model.MomentID)
	return model.ID, nil
}

// LikeDelete 删除
func (r *Repo) LikeDelete(ctx context.Context, userID, momentID uint32) error {
	if err := r.db.WithContext(ctx).Where("user_id=? and moment_id=?", userID, momentID).
		Delete(&model.MomentLikeModel{}).Error; err != nil {
		return errors.Wrapf(err, "[repo.moment_like] delete err")
	}
	r.delLikeCache(ctx, momentID)
	return nil
}

// LikeExist 记录是否存在
func (r *Repo) LikeExist(ctx context.Context, userID, momentID uint32) (is bool, err error) {
	userIds, err := r.GetLikeUserIdsByMomentID(ctx, momentID)
	if err != nil {
		return false, errors.Wrapf(err, "[repo.moment_like] userIds err")
	}
	return util.InuInt32Slice(userID, userIds), nil
}

// GetLikeUserIdsByMomentID 获取动态的所有点赞用户id列表
func (r *Repo) GetLikeUserIdsByMomentID(ctx context.Context, momentID uint32) (userIds []uint32, err error) {
	if err = r.queryCache(ctx, likeCacheKey(momentID), &userIds, func(data interface{}) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Model(&model.MomentLikeModel{}).Select("user_id").
			Where("moment_id=?", momentID).Order("id asc").Pluck("user_id", data).Error; err != nil {
			return errors.Wrap(err, "[repo.moment_like] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.moment_like] query cache")
	}
	return
}

// GetLikesByMomentIds 朋友圈动态点赞列表
func (r *Repo) GetLikesByMomentIds(ctx context.Context, ids []uint32) (likes map[uint32][]uint32, err error) {
	likes = make(map[uint32][]uint32, len(ids))

	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, likeCacheKey(id))
	}
	// 从cache批量获取
	cacheMap := make(map[string]*[]uint32)
	if err = r.cache.MultiGet(ctx, keys, cacheMap, func() interface{} {
		return &[]uint32{}
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.moment_like] multi get cache data err")
	}

	// 查询未命中
	for _, id := range ids {
		userIds, ok := cacheMap[likeCacheKey(id)]
		if !ok {
			uids, err := r.GetLikeUserIdsByMomentID(ctx, id)
			if err != nil {
				logger.Warnf("[repo.moment_like] get like err: %v", err)
				continue
			}
			userIds = &uids
		}
		if len(*userIds) == 0 {
			continue
		}
		likes[id] = *userIds
	}
	return likes, nil
}

//delLikeCache 删除缓存
func (r *Repo) delLikeCache(ctx context.Context, mid uint32) {
	if err := r.cache.Del(ctx, likeCacheKey(mid)); err != nil {
		logger.Warnf("[repo.moment_like] delete cache err:%v", err)
	}
}

func likeCacheKey(mid uint32) string {
	return fmt.Sprintf("moment:like:%d", mid)
}
