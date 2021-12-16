package repo

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"order/model"
)

//BatchCreateOrderItem 批量创建订单子项
func (r *Repo) BatchCreateOrderItem(ctx context.Context, tx *gorm.DB, items []*model.OrderItemModel) error {
	err := tx.WithContext(ctx).Model(&model.OrderItemModel{}).Create(&items).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.orderItem] batch create")
	}
	return nil
}