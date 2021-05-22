package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"chat/app/logic/model"
	"chat/pkg/cache"
	"chat/pkg/log"
	"chat/pkg/redis"
)

type IMomentTimeline interface {
	// 批量创建
	TimelineBatchCreate(ctx context.Context, tx *gorm.DB, models []*model.MomentTimelineModel) (ids []uint32, err error)
	// 记录是否存在
	TimelineExist(ctx context.Context, userId, momentId uint32) (bool, error)
}

// TimelineBatchCreate 批量创建
func (r *Repo) TimelineBatchCreate(ctx context.Context, tx *gorm.DB, models []*model.MomentTimelineModel) (ids []uint32, err error) {
	err = tx.Create(&models).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.moment_timeline] batch create err")
	}
	for _, tag := range models {
		ids = append(ids, tag.ID)
	}
	return ids, nil
}

// TimelineExist 记录是否存在
func (r *Repo) TimelineExist(ctx context.Context, userId, momentId uint32) (is bool, err error) {
	start := time.Now()
	defer func() {
		log.Infof("[repo.moment_timeline] get exist by uid:%d mid: %d cost: %d μs", userId, momentId, time.Since(start).Microseconds())
	}()
	var c int64
	// 从cache获取
	c, err = r.timelineCache.GetCache(ctx, userId, momentId)
	if err != nil {
		if err == cache.ErrPlaceholder {
			return false, nil
		} else if err != redis.Nil {
			return false, errors.Wrapf(err, "[repo.moment_timeline] get cache count by uid:%d mid: %d", userId, momentId)
		}
	}
	if c > 0 {
		log.Infof("[repo.moment_timeline] get count from cache, uid:%d mid: %d", userId, momentId)
		return true, nil
	}

	getDataFn := func() (interface{}, error) {
		// 从数据库中获取
		err = r.db.Model(&model.MomentTimelineModel{}).Where("user_id=? and moment_id=?", userId, momentId).Count(&c).Error
		if err != nil {
			return nil, errors.Wrapf(err, "[repo.moment_timeline] query db err")
		}

		// set cache
		err = r.timelineCache.SetCache(ctx, userId, momentId, c)
		if err != nil {
			return 0, errors.Wrap(err, "[repo.moment_timeline] set cache err")
		}
		return c, nil
	}

	g := singleflight.Group{}
	doKey := fmt.Sprintf("get_moment_timeline_count_%d_%d", userId, momentId)
	val, err, _ := g.Do(doKey, getDataFn)
	if err != nil {
		return false, errors.Wrap(err, "[repo.moment_timeline] get err via single flight do")
	}
	data := val.(int64)

	return data > 0, nil
}
