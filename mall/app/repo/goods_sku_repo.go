package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"mall/app/model"
)

//GetSkusByGID 商品SKU列表
func (r *Repo) GetSkusByGID(ctx context.Context, goodsID int) (list []*model.GoodsSkuModel, err error) {
	err = r.db.WithContext(ctx).Where("goods_id=?", goodsID).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.goodsSku] db find by goods_id: %v", goodsID)
	}
	return
}

//GetSkusAttrsByGID 商品销售属性列表
func (r *Repo) GetSkusAttrsByGID(ctx context.Context, goodsID int) (list []*model.GoodsSkuAttrModel, err error) {
	err = r.db.WithContext(ctx).Where("goods_id=?", goodsID).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.goodsSkuAttr[ db find by goods_id: %v", goodsID)
	}
	return
}

//GetSkuByID 获取SKU详情
func (r *Repo) GetSkuByID(ctx context.Context, id int) (*model.GoodsSkuModel, error) {
	sku := new(model.GoodsSkuModel)
	err := r.db.WithContext(ctx).First(&sku, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.goodsSku] db first")
	}
	return sku, nil
}