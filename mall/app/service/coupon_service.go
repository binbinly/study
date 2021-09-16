package service

import (
	"context"
	"github.com/pkg/errors"
	"mall/app/idl"
	"mall/app/model"
	"time"
)

var (
	//ErrCouponNotFound 优惠券不存在
	ErrCouponNotFound = errors.New("coupon not found")
	//ErrCouponNoNum 优惠券已领完
	ErrCouponNoNum = errors.New("coupon no num")
	//ErrCouponReceived 已领取过
	ErrCouponReceived = errors.New("coupon received")
)

//ICoupon 优惠券服务接口
type ICoupon interface {
	GetCouponList(ctx context.Context, userID, offset, limit int) ([]*model.Coupon, error)
	GetMyCouponList(ctx context.Context, userID, offset, limit int) ([]*model.Coupon, error)
	CouponDraw(ctx context.Context, userID, id int) error
}

//GetCouponList 优惠券列表
func (s *Service) GetCouponList(ctx context.Context, userID, offset, limit int) ([]*model.Coupon, error) {
	list, err := s.repo.GetCouponList(ctx, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, "[service.coupon] get list")
	}
	if len(list) == 0 {
		return []*model.Coupon{}, nil
	}
	//已领取的优惠券
	ids, err := s.repo.GetCouponUserDrawIds(ctx, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.coupon] get draw ids uid: %v", userID)
	}
	return idl.TransferCouponList(list, ids), nil
}

//GetMyCouponList 我的优惠券
func (s *Service) GetMyCouponList(ctx context.Context, userID, offset, limit int) ([]*model.Coupon, error) {
	list, err := s.repo.GetCouponUserList(ctx, userID, offset, limit)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.coupon] get my list")
	}
	return idl.TransferCouponList(list, nil), nil
}

//CouponDraw 领取优惠券
func (s *Service) CouponDraw(ctx context.Context, userID, id int) error {
	coupon, err := s.repo.GetCouponByID(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[service.coupon] find id: %v", id)
	}
	now := int(time.Now().Unix())
	if coupon.ID == 0 || coupon.StartAt > now || coupon.EndAt < now {
		return ErrCouponNotFound
	}
	exist, err := s.repo.CheckReceived(ctx, userID, id)
	if err != nil {
		return errors.Wrapf(err, "[service.coupon] check uid: %v id: %v", userID, id)
	}
	if exist {
		return ErrCouponReceived
	}
	couponUser, err := s.repo.GetFirstCouponUser(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[service.coupon] find coupon_user id: %v", id)
	}
	if couponUser.ID == 0 {
		return ErrCouponNoNum
	}
	couponUser.UserID = userID
	err = s.repo.CouponUserSave(ctx, couponUser)
	if err != nil {
		return errors.Wrapf(err, "[service.coupon] save")
	}
	return nil
}