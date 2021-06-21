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

//IMoment 朋友圈
type IMoment interface {
	// 创建一条动态
	MomentCreate(ctx context.Context, tx *gorm.DB, message *model.MomentModel) (id uint32, err error)
	// 我的朋友圈列表
	GetMyMoments(ctx context.Context, userID uint32, offset, limit int) ([]*model.MomentModel, error)
	// 指定好友的朋友圈
	GetMomentsByUserID(ctx context.Context, myID, userID uint32, offset, limit int) ([]*model.MomentModel, error)
	// 获取动态信息
	GetMomentByID(ctx context.Context, id uint32) (*model.MomentModel, error)
	// 批量获取动态信息
	GetMomentsByIds(ctx context.Context, ids []uint32) ([]*model.MomentModel, error)
}

// MomentCreate 创建
func (r *Repo) MomentCreate(ctx context.Context, tx *gorm.DB, moment *model.MomentModel) (id uint32, err error) {
	err = tx.WithContext(ctx).Create(&moment).Error
	if err != nil {
		return 0, errors.Wrap(err, "[repo.moment] create moment err")
	}
	return moment.ID, nil
}

// GetMyMoments 我的朋友圈列表
func (r *Repo) GetMyMoments(ctx context.Context, userID uint32, offset, limit int) (list []*model.MomentModel, err error) {
	var ids []uint32
	err = r.db.WithContext(ctx).Model(&model.MomentLike{}).Raw("select moment_id from moment_timeline where user_id = ? order by id desc limit ? offset ?", userID, limit, offset).Pluck("moment_id", &ids).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.moment] get db ids err")
	}
	return r.GetMomentsByIds(ctx, ids)
}

// GetMomentsByUserID 指定用户的朋友圈
func (r *Repo) GetMomentsByUserID(ctx context.Context, myID, userID uint32, offset, limit int) (list []*model.MomentModel, err error) {
	if myID == userID { // 查看自己
		err = r.db.WithContext(ctx).Raw("select * from moment where user_id=? order by id desc limit ? offset ?", myID, limit, offset).Find(&list).Error
	} else {
		err = r.db.WithContext(ctx).Raw("select * from moment where user_id=? and (see_type=1 or (see_type = 3 and FIND_IN_SET(?,see)) or (see_type = 4 and !FIND_IN_SET(?,see)) ) order by id desc limit ? offset ?", userID, myID, myID, limit, offset).Find(&list).Error
	}
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.moment] list err")
	}
	return list, nil
}

// GetMomentByID 获取动态信息
func (r *Repo) GetMomentByID(ctx context.Context, id uint32) (moment *model.MomentModel, err error) {
	start := time.Now()
	defer func() {
		log.Debugf("[repo.moment] get moment by id: %d cost: %d μs", id, time.Since(start).Microseconds())
	}()
	// 从cache获取
	moment, err = r.momentCache.GetCache(ctx, id)
	if err != nil {
		if err == cache.ErrPlaceholder {
			return new(model.MomentModel), nil
		} else if err != redis.Nil {
			// fail fast, if cache error return, don't request to db
			return nil, errors.Wrapf(err, "[repo.moment] get moment by id: %d", id)
		}
	}
	// hit cache
	if moment != nil {
		log.Debugf("[repo.moment] get moment data from cache, id: %d", id)
		return
	}

	getDataFn := func() (interface{}, error) {
		data := new(model.MomentModel)
		// 从数据库中获取
		err = r.db.WithContext(ctx).First(data, id).Error
		// if data is empty, set not found cache to prevent cache penetration(缓存穿透)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = r.momentCache.SetCacheWithNotFound(ctx, id)
			if err != nil {
				log.Warnf("[repo.moment] SetCacheWithNotFound err, id: %d", id)
			}
			return data, nil
		} else if err != nil {
			return nil, errors.Wrapf(err, "[repo.moment] query db err")
		}

		// set cache
		err = r.momentCache.SetCache(ctx, id, data)
		if err != nil {
			return data, errors.Wrap(err, "[repo.moment] set cache data err")
		}
		return data, nil
	}

	g := singleflight.Group{}
	doKey := fmt.Sprintf("get_moment_%d", id)
	val, err, _ := g.Do(doKey, getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.moment] get moment err via single flight do")
	}
	data := val.(*model.MomentModel)

	return data, nil
}

// 批量获取动态信息
func (r *Repo) GetMomentsByIds(ctx context.Context, ids []uint32) (moments []*model.MomentModel, err error) {
	// 从cache批量获取
	cacheMap, err := r.momentCache.MultiGetCache(ctx, ids)
	if err != nil {
		return moments, errors.Wrap(err, "[repo.moment] multi get moment cache data err")
	}

	// 查询未命中
	for _, id := range ids {
		idx := r.momentCache.GetCacheKey(id)
		moment, ok := cacheMap[idx]
		if !ok {
			moment, err = r.GetMomentByID(ctx, id)
			if err != nil {
				log.Warnf("[repo.moment] get moment model err: %v", err)
				continue
			}
		}
		moments = append(moments, moment)
	}
	return moments, nil
}
