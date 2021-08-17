package repository

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"

	"chat/app/chat/model"
	"chat/pkg/cache"
	"chat/pkg/redis"
)

//IEmoticon 表情包仓库接口
type IEmoticon interface {
	GetEmoticonCatAll(ctx context.Context) (list []*model.Emoticon, err error)
	GetEmoticonListByCat(ctx context.Context, cat string) (list []*model.Emoticon, err error)
}

//GetEmoticonCatAll 获取表情所有分类
func (r *Repo) GetEmoticonCatAll(ctx context.Context) (list []*model.Emoticon, err error) {
	// 从cache获取
	list, err = r.emoCache.GetCatCache(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.emoticon] get cat all from cache")
	}
	// hit cache
	if len(list) > 0 {
		return
	}

	getDataFn := func() (interface{}, error) {
		data := make([]*model.Emoticon, 0)
		// 从数据库中获取
		err = r.db.WithContext(ctx).Model(&model.EmoticonModel{}).Group("category").Scan(&data).Error
		// if data is empty, set not found cache to prevent cache penetration(缓存穿透)
		if err != nil {
			return nil, errors.Wrap(err, "[repo.emoticon] query db err")
		}

		// set cache
		err = r.emoCache.SetCatCache(ctx, data)
		if err != nil {
			return data, errors.Wrap(err, "[repo.emoticon] set cache data err")
		}
		return data, nil
	}

	gr := singleflight.Group{}
	val, err, _ := gr.Do("emoticon_cat_all", getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.emoticon] cat all err via single flight do")
	}
	data := val.([]*model.Emoticon)
	return data, nil
}

//GetEmoticonListByCat 获取分类下所有表情
func (r *Repo) GetEmoticonListByCat(ctx context.Context, cat string) (list []*model.Emoticon, err error) {
	// 从cache获取
	list, err = r.emoCache.GetCache(ctx, cat)
	if err != nil {
		if err == cache.ErrPlaceholder {
			return make([]*model.Emoticon, 0), nil
		} else if err != redis.Nil {
			// fail fast, if cache error return, don't request to db
			return nil, errors.Wrapf(err, "[repo.emoticon] get cache by cat: %s", cat)
		}
	}
	// hit cache
	if len(list) > 0 {
		return
	}

	getDataFn := func() (interface{}, error) {
		data := make([]*model.Emoticon, 0)
		// 从数据库中获取
		err = r.db.WithContext(ctx).Model(&model.EmoticonModel{}).Where("category=?", cat).Scan(&data).Error
		// if data is empty, set not found cache to prevent cache penetration(缓存穿透)
		if err != nil {
			return nil, errors.Wrapf(err, "[repo.emoticon] query db err from cat:%s", cat)
		}

		// set cache
		err = r.emoCache.SetCache(ctx, cat, data)
		if err != nil {
			return data, errors.Wrapf(err, "[repo.emoticon] set cache data err from cat:%s", cat)
		}
		return data, nil
	}

	gr := singleflight.Group{}
	val, err, _ := gr.Do(fmt.Sprintf("get_emoticon_%s", cat), getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.emoticon] get err via single flight do")
	}
	data := val.([]*model.Emoticon)
	return data, nil
}
