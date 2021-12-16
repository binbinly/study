package repo

import (
	"context"
	"fmt"
	"go-micro.dev/v4/logger"
	"reflect"

	"github.com/pkg/errors"
	"product/model"
)

//GetBrandByID 获取品牌
func (r *Repo) GetBrandByID(ctx context.Context, id int64) (brand *model.Brand, err error) {
	if err = r.QueryCache(ctx, buildBrandCacheKey(id), &brand, func(data interface{}) error {
		// 从数据库中获取
		brand := &model.Brand{}
		if err := r.DB.WithContext(ctx).Model(&model.BrandModel{}).Where("id=?", id).Scan(&brand).Error; err != nil {
			return errors.Wrapf(err, "[repo.brand] query db")
		}
		reflect.ValueOf(data).Elem().Set(reflect.ValueOf(brand))
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.brand] query cache")
	}

	return
}

//GetBrandsByIds 批量获取品牌列表
func (r *Repo) GetBrandsByIds(ctx context.Context, ids []int64) (brands map[int64]*model.Brand, err error) {
	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, buildBrandCacheKey(id))
	}
	// 从cache批量获取
	cacheMap := make(map[string]*model.Brand)
	if err = r.Cache.MultiGetCache(ctx, keys, cacheMap, func() interface{} {
		return &model.Brand{}
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.brand] multi get brand cache data err")
	}

	brands = make(map[int64]*model.Brand, len(ids))
	// 查询未命中
	for _, id := range ids {
		brand, ok := cacheMap[buildBrandCacheKey(id)]
		if !ok {
			brand, err = r.GetBrandByID(ctx, id)
			if err != nil {
				logger.Warnf("[repo.brand] get brand model err: %v", err)
				continue
			}
			if brand.ID == 0 {
				continue
			}
		}
		brands[id] = brand
	}
	return brands, nil
}

//buildBrandCacheKey 构建品牌模型缓存键
func buildBrandCacheKey(id int64) string {
	return fmt.Sprintf("brand:%d", id)
}