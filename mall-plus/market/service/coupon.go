package service

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"common/errno"
	"common/orm"
	pb "common/proto/market"
	"common/util"
	"market/idl"
	"market/model"
	"pkg/redis"
)

//ICoupon 优惠券接口
type ICoupon interface {
	GetCouponList(ctx context.Context, memberID, skuID int64) ([]*pb.Coupon, error)
	GetMyCouponList(ctx context.Context, memberID int64) ([]*pb.Coupon, error)
	CouponDraw(ctx context.Context, memberID, id int64) error
	CouponUsed(ctx context.Context, memberID, id, orderID int64) error
	GetCouponInfo(ctx context.Context, memberID, id int64) (*pb.Coupon, error)
}

//GetCouponList 优惠券列表
func (s *Service) GetCouponList(ctx context.Context, memberID, skuID int64) ([]*pb.Coupon, error) {
	//TODO 获取sku详情
	list, err := s.repo.GetCouponList(ctx, 0, 0)
	if err != nil {
		return nil, errors.Wrap(err, "[service.coupon] get list")
	}
	if len(list) == 0 {
		return []*pb.Coupon{}, nil
	}
	//已领取的优惠券
	ids, err := s.repo.GetCouponUserDrawIds(ctx, memberID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.coupon] get draw ids uid: %v", memberID)
	}
	return idl.TransferCouponList(list, ids), nil
}

//GetMyCouponList 我的优惠券
func (s *Service) GetMyCouponList(ctx context.Context, memberID int64) ([]*pb.Coupon, error) {
	list, err := s.repo.GetCouponMemberList(ctx, memberID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.coupon] get my list")
	}
	return idl.TransferMyCouponList(list), nil
}

//CouponDraw 领取优惠券
func (s *Service) CouponDraw(ctx context.Context, memberID, id int64) error {
	coupon, err := s.repo.GetCouponByID(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[service.coupon] find id: %v", id)
	}
	now := time.Now().Unix()
	if coupon == nil || coupon.ID == 0 || coupon.StartAt > now || coupon.EndAt < now {
		return errno.ErrCouponNotFound
	}
	exist, err := s.repo.CheckReceived(ctx, memberID, id)
	if err != nil {
		return errors.Wrapf(err, "[service.coupon] check uid: %v id: %v", memberID, id)
	}
	if exist {
		return errno.ErrCouponReceived
	}
	num, err := redis.Client.Decr(ctx, fmt.Sprintf("coupon_num:%d", id)).Result()
	if err == redis.Nil {
		return errno.ErrCouponFinished
	} else if err != nil {
		return errors.Wrapf(err, "[service.coupon] redis decr")
	}
	if num < 0 { //已领取完
		return errno.ErrCouponFinished
	}
	couponMember := &model.CouponMemberModel{
		MID:      orm.MID{MemberID: memberID},
		CouponID: id,
		GetType:  model.CouponGetTypeDraw,
		Status:   model.CouponStatusInit,
	}
	err = s.repo.CouponUserSave(ctx, couponMember)
	if err != nil {
		return errors.Wrapf(err, "[service.coupon] save")
	}
	return nil
}

//CouponUsed 使用优惠券
func (s *Service) CouponUsed(ctx context.Context, memberID, id, orderID int64) error {
	mCoupon, err := s.repo.GetCouponMemberByID(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[service.coupon] get coupon user by id: %v", id)
	}
	if mCoupon.ID == 0 || mCoupon.MemberID != memberID || mCoupon.Status != model.CouponStatusInit {
		return errno.ErrCouponNotFound
	}
	err = s.repo.SetCouponMemberUsed(ctx, id, memberID, orderID)
	if errors.Is(err, util.ErrNotRecordUpdate) {
		return errno.ErrCouponNotFound
	} else if err != nil {
		return errors.Wrapf(err, "[service.coupon] set used id: %v uid: %v, oid: %v", id, memberID, orderID)
	}
	return nil
}

//GetCouponInfo 获取优惠券详情
func (s *Service) GetCouponInfo(ctx context.Context, memberID, id int64) (*pb.Coupon, error) {
	mCoupon, err := s.repo.GetCouponMemberByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.coupon] get coupon user by id: %v", id)
	}
	if mCoupon.ID == 0 || mCoupon.MemberID != memberID || mCoupon.Status != model.CouponStatusInit {
		return nil, errno.ErrCouponNotFound
	}
	coupon, err := s.repo.GetCouponByID(ctx, mCoupon.CouponID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.coupon] get coupon by id: %v", mCoupon.CouponID)
	}
	now := time.Now().Unix()
	if coupon == nil || coupon.ID == 0 || now > coupon.EndAt || now < coupon.StartAt {
		return nil, errno.ErrCouponNotFound
	}

	res := idl.TransferCouponInfo(coupon)
	res.Status = int32(mCoupon.Status)
	return res, nil
}
