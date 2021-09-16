package repo

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"mall/app/model"
)

//GetCouponUserByID 用户领取的优惠券详情
func (r *Repo) GetCouponUserByID(ctx context.Context, id int) (*model.CouponUserModel, error) {
	coupon := new(model.CouponUserModel)
	err := r.db.WithContext(ctx).First(coupon, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.coupon_user] first db id: %v", id)
	}
	return coupon, nil
}

//GetFirstCouponUser 获取第一条可被领取记录
func (r *Repo) GetFirstCouponUser(ctx context.Context, id int) (*model.CouponUserModel, error) {
	coupon := new(model.CouponUserModel)
	err := r.db.WithContext(ctx).Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("coupon_id=? && user_id=0", id).First(&coupon).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.coupon_user] db by find id: %v", id)
	}
	return coupon, nil
}

//SetCouponUserUsed 设置优惠券已使用
func (r *Repo) SetCouponUserUsed(ctx context.Context, tx *gorm.DB, id, userID int) error {
	err := tx.WithContext(ctx).Model(&model.CouponUserModel{}).Where("id=? and user_id=? and is_used=0", id, userID).Update("is_used", 1).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.coupon_user] set used")
	}
	return nil
}

//CouponUserSave 保存记录
func (r *Repo) CouponUserSave(ctx context.Context, coupon *model.CouponUserModel) error {
	err := r.db.WithContext(ctx).Save(coupon).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.coupon_user] save db")
	}
	return nil
}

//GetCouponUserDrawIds 已领取的优惠券id列表
func (r *Repo) GetCouponUserDrawIds(ctx context.Context, userID int) (ids []int, err error) {
	err = r.db.WithContext(ctx).Model(&model.CouponUserModel{}).Where("user_id=?", userID).Pluck("coupon_id", &ids).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.coupon_user] draw pluck db")
	}
	return
}

//CheckReceived 检查是否已领取过
func (r *Repo) CheckReceived(ctx context.Context, userID, id int) (bool, error) {
	var c int64
	err := r.db.WithContext(ctx).Model(&model.CouponUserModel{}).Where("user_id=? && coupon_id=?",userID, id).Count(&c).Error
	if err != nil {
		return false, errors.Wrapf(err, "[repo.coupon_user] count")
	}
	return c > 0, nil
}