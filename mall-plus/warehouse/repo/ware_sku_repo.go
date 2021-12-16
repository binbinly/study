package repo

import (
	"context"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"go-micro.dev/v4/logger"
	"gorm.io/gorm"

	"warehouse/model"
)

//GetWareSkuStock 获取sku库存数量
func (r *Repo) GetWareSkuStock(ctx context.Context, skuID int64) (info *model.WareSkuStock, err error) {
	if err = r.QueryCache(ctx, buildCacheKey(skuID), &info, func(data interface{}) error {
		// 从数据库中获取
		stock := &model.WareSkuStock{}
		if err := r.DB.WithContext(ctx).Model(&model.WareSkuModel{}).Select("sku_id,sum(stock) as stock,sum(stock_lock) as stock_lock").
			Where("sku_id=?", skuID).Group("sku_id").Scan(&stock).Error; err != nil {
			return errors.Wrapf(err, "[repo.wareSku] ware sku stock query db")
		}
		reflect.ValueOf(data).Elem().Set(reflect.ValueOf(stock))
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.wareSku] query cache")
	}

	return
}

//BatchGetWareSkuStocks 批量获取spu下所有sku库存数量
func (r *Repo) BatchGetWareSkuStocks(ctx context.Context, spuID int64, skuIds []int64) (list []*model.WareSkuStock, err error) {
	doKey := fmt.Sprintf("spu_stock:%d", spuID)
	if err = r.QueryCache(ctx, doKey, &list, func(data interface{}) error {
		// 从数据库中获取
		if err := r.DB.WithContext(ctx).Model(&model.WareSkuModel{}).Select("sku_id,sum(stock) as stock,sum(stock_lock) as stock_lock").
			Where("sku_id in ?", skuIds).Group("sku_id").Scan(data).Error; err != nil {
			return errors.Wrapf(err, "[repo.wareSku] batch ware sku stock query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.wareSku] query cache")
	}

	return
}

//BatchGetWareSkus 批量获取sku库存信息
func (r *Repo) BatchGetWareSkus(ctx context.Context, skuIds []int64) (list []*model.WareSkuModel, err error) {
	if err = r.DB.WithContext(ctx).Model(&model.WareSkuModel{}).Where("sku_id in ?", skuIds).Find(&list).Error; err != nil {
		return nil, errors.Wrap(err, "[repo.wareSku] batch ware skus")
	}
	return list, nil
}

//WareSkuSave 更新库存信息
func (r *Repo) WareSkuSave(ctx context.Context, tx *gorm.DB, ware *model.WareSkuModel) error {
	if err := tx.WithContext(ctx).Save(ware).Error; err != nil {
		return errors.Wrap(err, "[repo.wareSku] save")
	}
	//清除缓存
	if err := r.Cache.DelCache(ctx, buildCacheKey(ware.SkuID)); err != nil {
		logger.Warnf("[repo.wareSku] del cache key: %v", buildCacheKey(ware.SkuID))
	}
	return nil
}

//buildCacheKey 构建缓存键
func buildCacheKey(skuID int64) string {
	return fmt.Sprintf("sku_stock:%d", skuID)
}
