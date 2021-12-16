package repo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"common/orm"
	"product/model"
)

//GetImagesBySkuID 获取sku图集
func (r *Repo) GetImagesBySkuID(ctx context.Context, skuID int64) (list []*model.SkuImageModel, err error) {
	doKey := fmt.Sprintf("sku_image:%d", skuID)
	if err = r.QueryCache(ctx, doKey, &list, func(data interface{}) error {
		// 从数据库中获取
		if err := r.DB.WithContext(ctx).Model(&model.SkuImageModel{}).Where("sku_id=?", skuID).
			Order(orm.DefaultOrderSort).Find(&list).Error; err != nil {
			return errors.Wrap(err, "[repo.skuImage] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.skuImage] query cache")
	}
	return
}
