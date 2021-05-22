package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"

	"chat/app/logic/model"
	"chat/pkg/cache"
	"chat/pkg/log"
	"chat/pkg/redis"
)

type IMomentLike interface {
	// 创建
	LikeCreate(ctx context.Context, model *model.MomentLikeModel) (id uint32, err error)
	// 删除数据
	LikeDelete(ctx context.Context, user, moment uint32) error
	// 是否存在
	LikeExist(ctx context.Context, userId, momentId uint32) (bool, error)
	// 朋友圈动态点赞用户列表
	GetLikeUserIdsByMomentId(ctx context.Context, momentId uint32) (*[]uint32, error)
	// 朋友圈动态点赞列表
	GetLikesByMomentIds(ctx context.Context, momentIds []uint32) (map[uint32]*[]uint32, error)
}

// LikeCreate 创建
func (r *Repo) LikeCreate(ctx context.Context, model *model.MomentLikeModel) (id uint32, err error) {
	err = r.db.Create(model).Error
	if err != nil {
		return 0, errors.Wrapf(err, "[repo.moment_like] create err")
	}
	// 删除cache
	err = r.likeCache.DelCache(ctx, model.MomentId)
	if err != nil {
		log.Warnf("[repo.moment_like] delete cache err:%v", err)
	}
	return model.ID, nil
}

// LikeDelete 删除
func (r *Repo) LikeDelete(ctx context.Context, userId, momentId uint32) error {
	err := r.db.Where("user_id=? and moment_id=?", userId, momentId).Delete(&model.MomentLikeModel{}).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.moment_like] delete err")
	}
	// 删除cache
	err = r.likeCache.DelCache(ctx, momentId)
	if err != nil {
		log.Warnf("[repo.moment_like] delete cache err:%v", err)
	}
	return nil
}

// LikeExist 记录是否存在
func (r *Repo) LikeExist(ctx context.Context, userId, momentId uint32) (is bool, err error) {
	userIds, err := r.GetLikeUserIdsByMomentId(ctx, momentId)
	if err != nil {
		return false, errors.Wrapf(err, "[repo.moment_like] userIds err")
	}
	for _, uid := range *userIds {
		if uid == userId {
			return true, nil
		}
	}
	return false, nil
}

// GetLikeUserIdsByMomentId 获取动态的所有点赞用户id列表
func (r *Repo) GetLikeUserIdsByMomentId(ctx context.Context, momentId uint32) (userIds *[]uint32, err error) {
	start := time.Now()
	defer func() {
		log.Infof("[repo.moment_like] get userIds by mid: %d cost: %d μs", momentId, time.Since(start).Microseconds())
	}()
	// 从cache获取
	userIds, err = r.likeCache.GetCache(ctx, momentId)
	if err != nil {
		if err == cache.ErrPlaceholder {
			return new([]uint32), nil
		} else if err != redis.Nil {
			return nil, errors.Wrapf(err, "[repo.moment_like] get cache userIds by mid: %d", momentId)
		}
	}
	if userIds != nil {
		log.Infof("[repo.moment_like] get userIds from cache, mid: %d", momentId)
		return
	}

	getDataFn := func() (interface{}, error) {
		data := make([]uint32, 0)
		// 从数据库中获取
		err = r.db.Model(&model.MomentLikeModel{}).Select("user_id").Where("moment_id=?", momentId).Order("id asc").Pluck("user_id", &data).Error
		if err != nil {
			return nil, errors.Wrapf(err, "[repo.moment_like] query db err")
		}

		// set cache
		err = r.likeCache.SetCache(ctx, momentId, data)
		if err != nil {
			return data, errors.Wrap(err, "[repo.moment_like] set cache err")
		}
		return &data, nil
	}

	g := singleflight.Group{}
	doKey := fmt.Sprintf("get_moment_like_%d", momentId)
	val, err, _ := g.Do(doKey, getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.moment_like] get err via single flight do")
	}
	data := val.(*[]uint32)

	return data, nil
}

// GetLikesByMomentIds 朋友圈动态点赞列表
func (r *Repo) GetLikesByMomentIds(ctx context.Context, momentIds []uint32) (likes map[uint32]*[]uint32, err error) {
	likes = make(map[uint32]*[]uint32, 0)

	// 从cache批量获取
	likeMap, err := r.likeCache.MultiGetCache(ctx, momentIds)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.moment_like] multi get likes cache data err")
	}

	// 查询未命中
	for _, momentId := range momentIds {
		idx := r.likeCache.GetCacheKey(momentId)
		userIds, ok := likeMap[idx]
		if !ok {
			userIds, err = r.GetLikeUserIdsByMomentId(ctx, momentId)
			if err != nil {
				log.Warnf("[repo.moment_like] get userIds err: %v", err)
				continue
			}
		}
		likes[momentId] = userIds
	}
	return likes, nil
}
