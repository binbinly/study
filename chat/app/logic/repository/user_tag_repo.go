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
	"chat/pkg/utils"
)

//IUserTag 用户标签
type IUserTag interface {
	// 用户所有标签
	GetTagsByUserID(ctx context.Context, userID uint32) (list []*model.UserTag, err error)
	// 获取对应标签名
	GetTagNamesByIds(ctx context.Context, userID uint32, ids []uint32) (names []string, err error)
	// 批量创建
	TagBatchCreate(ctx context.Context, tags []*model.UserTagModel) (ids []uint32, err error)
}

// GetTagsByUserID 获取用户收藏列表
func (r *Repo) GetTagsByUserID(ctx context.Context, userID uint32) (list []*model.UserTag, err error) {
	start := time.Now()
	defer func() {
		log.Debugf("[repo.tag] uid: %d cost: %d μs", userID, time.Since(start).Microseconds())
	}()
	// 从cache获取
	list, err = r.tagCache.GetCacheAll(ctx, userID)
	if err != nil {
		if err == cache.ErrPlaceholder {
			return make([]*model.UserTag, 0), nil
		} else if err != redis.Nil {
			// fail fast, if cache error return, don't request to db
			return nil, errors.Wrapf(err, "[repo.tag] get by uid: %d", userID)
		}
	}
	if len(list) > 0 {
		log.Debugf("[repo.tag] get from cache, uid: %d", userID)
		return
	}

	getDataFn := func() (interface{}, error) {
		data := make([]*model.UserTag, 0)
		err = r.db.WithContext(ctx).Model(&model.UserTagModel{}).Where("user_id = ? ", userID).Order(model.DefaultOrder).Scan(&data).Error
		if err != nil {
			return nil, errors.Wrapf(err, "[repo.tag] query db err")
		}

		// set cache
		err = r.tagCache.SetCacheAll(ctx, userID, data)
		if err != nil {
			return data, errors.Wrap(err, "[repo.tag] set cache all err")
		}
		return data, nil
	}

	gr := singleflight.Group{}
	doKey := fmt.Sprintf("get_tag_all_%d", userID)
	val, err, _ := gr.Do(doKey, getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.tag] get all err via single flight do")
	}
	data := val.([]*model.UserTag)

	return data, nil
}

// GetTagNamesByIds 标签id获取标签名列表
func (r *Repo) GetTagNamesByIds(ctx context.Context, userID uint32, ids []uint32) (names []string, err error) {
	tags, err := r.GetTagsByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.tag] get all err")
	}
	names = make([]string, 0)
	for _, tag := range tags {
		if utils.InuInt32Slice(tag.ID, ids) {
			names = append(names, tag.Name)
		}
	}
	return
}

// TagBatchCreate 批量创建标签
func (r *Repo) TagBatchCreate(ctx context.Context, tags []*model.UserTagModel) (ids []uint32, err error) {
	err = r.db.WithContext(ctx).Create(&tags).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.tag] batch create err")
	}
	// 删除缓存
	err = r.tagCache.DelCacheAll(ctx, tags[0].UserID)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.tag] delete all cache")
	}
	for _, tag := range tags {
		ids = append(ids, tag.ID)
	}
	return ids, nil
}
