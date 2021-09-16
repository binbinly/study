package repo

import (
	"context"
	"github.com/pkg/errors"
	"mall/app/model"
)

//GetImagesByGID 获取商品图片
func (r *Repo) GetImagesByGID(ctx context.Context, goodsID int) (list []*model.GoodsImageModel, err error) {
	err = r.db.WithContext(ctx).Where("goods_id=?",goodsID).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.goodsImage] db find by goods_id: %v", goodsID)
	}
	return list, nil
}