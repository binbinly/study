package idl

import (
	pb "common/proto/market"
	"common/util"
	"market/model"
	"pkg/utils"
)

//TransferCouponList 转换优惠券列表输出
func TransferCouponList(list []*model.CouponModel, ids []int64) []*pb.Coupon {
	if len(list) == 0 {
		return []*pb.Coupon{}
	}
	res := make([]*pb.Coupon, 0, len(list))
	for _, c := range list {
		status := model.CouponStatusNotReceive
		if utils.InInt64Slice(c.ID, ids) {
			status = model.CouponStatusInit
		}
		res = append(res, &pb.Coupon{
			Id:       c.ID,
			Name:     c.Name,
			Amount:   util.ParseAmount(c.Amount),
			MinPoint: util.ParseAmount(c.MinPoint),
			StartAt:  c.StartAt,
			EndAt:    c.EndAt,
			Note:     c.Note,
			Status:   int32(status),
		})
	}
	return res
}

//TransferMyCouponList 转换优惠券列表输出
func TransferMyCouponList(list []*model.Coupon) []*pb.Coupon {
	if len(list) == 0 {
		return []*pb.Coupon{}
	}
	res := make([]*pb.Coupon, 0, len(list))
	for _, c := range list {
		res = append(res, &pb.Coupon{
			Id:       c.ID,
			Name:     c.Name,
			Amount:   util.ParseAmount(c.Amount),
			MinPoint: util.ParseAmount(c.MinPoint),
			StartAt:  c.StartAt,
			EndAt:    c.EndAt,
			Note:     c.Note,
			Status:   int32(c.Status),
		})
	}
	return res
}

//TransferCouponInfo 转换优惠券详情
func TransferCouponInfo(c *model.CouponModel) *pb.Coupon {
	return &pb.Coupon{
		Id:       c.ID,
		Name:     c.Name,
		Amount:   util.ParseAmount(c.Amount),
		MinPoint: util.ParseAmount(c.MinPoint),
		StartAt:  c.StartAt,
		EndAt:    c.EndAt,
		Note:     c.Note,
	}
}
