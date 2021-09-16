package idl

import (
	"mall/app/constvar"
	"mall/app/model"
	"mall/pkg/utils"
)

//TransferCouponList 转换优惠券列表输出
func TransferCouponList(list []*model.CouponModel, ids []int) []*model.Coupon {
	res := make([]*model.Coupon, 0, len(list))
	for _, c := range list {
		status := model.CouponStatusNormal
		if utils.InIntSlice(c.ID, ids) {
			status = model.CouponStatusReceived
		}
		res = append(res, &model.Coupon{
			ID:       c.ID,
			Name:     c.Name,
			Type:     c.Type,
			Value:    constvar.ParseAmount(c.Value),
			MinPrice: constvar.ParseAmount(c.MinPrice),
			StartAt:  c.StartAt,
			EndAt:    c.EndAt,
			Desc:     c.Desc,
			Status:   status,
		})
	}
	return res
}
