package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"mall/app/model"
)

//CreateOrderAddress 创建订单收货地址
func (r *Repo) CreateOrderAddress(ctx context.Context, tx *gorm.DB, address *model.OrderAddressModel) error {
	err := tx.WithContext(ctx).Create(address).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.orderAddress] create")
	}
	return nil
}