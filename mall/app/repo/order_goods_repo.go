package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"mall/app/model"
)

//CreateOrderGoods 批量创建订单商品
func (r *Repo) CreateOrderGoods(ctx context.Context, tx *gorm.DB, goods []*model.OrderGoodsModel) error {
	err := tx.WithContext(ctx).Model(&model.OrderGoodsModel{}).Create(&goods).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.orderGoods] batch create")
	}
	return nil
}