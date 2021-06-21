package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"

	"chat/app/logic/model"
	"chat/pkg/cache"
	"chat/pkg/log"
	"chat/pkg/redis"
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
	err = r.db.WithContext(ctx).Create(collect).Error
	if err != nil {
		return 0, errors.Wrap(err, "[repo.collect] create collect err")
	}

	// 删除列表缓存
	err = r.collectCache.DelCacheList(ctx, collect.UserID)
	if err != nil {
		return 0, errors.Wrap(err, "[repo.apply] delete list cache")
	}
	return collect.ID, nil
}

// GetCollectsByUserID 获取用户收藏列表
func (r *Repo) GetCollectsByUserID(ctx context.Context, userID uint32, offset, limit int) (list []*model.CollectModel, err error) {
	start := time.Now()
	defer func() {
		log.Debugf("[repo.collect] collect list uid: %d offset: %d cost: %d μs", userID, offset, time.Since(start).Microseconds())
	}()
	// 从cache获取
	list, err = r.collectCache.GetCacheList(ctx, userID, strconv.Itoa(offset))
	if err != nil {
		if err == cache.ErrPlaceholder {
			return make([]*model.CollectModel, 0), nil
		} else if err != redis.Nil {
			return nil, errors.Wrapf(err, "[repo.collect] cache collect list by uid: %d offset:%d", userID, offset)
		}
	}
	if len(list) > 0 {
		log.Debugf("[repo.collect] get collect list from cache, uid: %d offset: %d", userID, offset)
		return
	}

	getDataFn := func() (interface{}, error) {
		data := make([]*model.CollectModel, 0)
		err = r.db.WithContext(ctx).Scopes(model.OffsetPage(offset, limit)).Where("user_id = ? ", userID).
			Order(model.DefaultOrder).Find(&data).Error
		if err != nil {
			return nil, errors.Wrap(err, "[repo.collect] get collect list err")
		}

		// set cache
		err = r.collectCache.SetCacheList(ctx, userID, strconv.Itoa(offset), data)
		if err != nil {
			return 0, errors.Wrap(err, "[repo.collect] set cache list err")
		}
		return data, nil
	}

	g := singleflight.Group{}
	doKey := fmt.Sprintf("get_collect_list_%d_%d", userID, offset)
	val, err, _ := g.Do(doKey, getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.collect] get collect list err via single flight do")
	}
	data := val.([]*model.CollectModel)
	return data, nil
}

// CollectDelete 删除收藏
func (r *Repo) CollectDelete(ctx context.Context, userID, id uint32) (err error) {
	err = r.db.WithContext(ctx).Where("user_id=?", userID).Delete(&model.CollectModel{}, id).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.collect] destroy err")
	}

	// 删除列表缓存
	err = r.collectCache.DelCacheList(ctx, userID)
	if err != nil {
		return errors.Wrap(err, "[repo.collect] delete list cache")
	}
	return nil
}
