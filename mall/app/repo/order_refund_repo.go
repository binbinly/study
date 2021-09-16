package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"mall/app/model"
)

//CreateOrderRefund 创建订单退款记录
func (r *Repo) CreateOrderRefund(ctx context.Context, tx *gorm.DB, refund *model.OrderRefundModel) error {
	err := tx.WithContext(ctx).Model(&model.OrderRefundModel{}).Create(&refund).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.orderRefund] create")
	}
	return nil
}