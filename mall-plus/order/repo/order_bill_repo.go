package repo

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"order/model"
)

//CreateOrderBill 创建发票
func (r *Repo) CreateOrderBill(ctx context.Context, tx *gorm.DB, bill *model.OrderBillModel) error {
	err := tx.WithContext(ctx).Create(bill).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.orderBill] create")
	}
	return nil
}