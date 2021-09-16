package repo

import (
	"context"
	"github.com/pkg/errors"
	"mall/app/model"
)

//GetAttrByGID 商品属性参数列表
func (r *Repo) GetAttrByGID(ctx context.Context, goodsID int) (list []*model.GoodsAttrModel, err error) {
	err = r.db.WithContext(ctx).Where("goods_id=?",goodsID).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.goodsAttr] db find by goods_id: %v", goodsID)
	}
	return
}