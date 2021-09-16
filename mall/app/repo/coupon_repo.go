package repo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"mall/app/model"
)

//ICoupon 优惠券接口
type ICoupon interface {
	GetCouponList(ctx context.Context, offset, limit int) (list []*model.CouponModel, err error)
	GetCouponUserList(ctx context.Context, userID, offset, limit int) (list []*model.CouponModel, err error)
	GetCouponByID(ctx context.Context, id int) (*model.CouponModel, error)

	GetCouponUserByID(ctx context.Context, id int) (*model.CouponUserModel, error)
	GetFirstCouponUser(ctx context.Context, id int) (*model.CouponUserModel, error)
	SetCouponUserUsed(ctx context.Context, tx *gorm.DB, id, userID int) error
	GetCouponUserDrawIds(ctx context.Context, userID int) (ids []int, err error)
	CheckReceived(ctx context.Context, userID, id int) (bool, error)
	CouponUserSave(ctx context.Context, coupon *model.CouponUserModel) error
}

//GetCouponList 有效优惠券列表
func (r *Repo) GetCouponList(ctx context.Context, offset, limit int) (list []*model.CouponModel, err error) {
	now := time.Now().Unix()
	err = r.db.WithContext(ctx).Model(&model.CouponModel{}).Scopes(model.OffsetPage(offset, limit), model.WhereRelease).
		Where("start_at<=?",now).Where("end_at>=?",now).
		Order("sort desc,id desc").Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.coupon] by db list")
	}
	return
}

//GetCouponUserList 获取用户领取的优惠券列表
func (r *Repo) GetCouponUserList(ctx context.Context, userID, offset, limit int) (list []*model.CouponModel, err error) {
	err = r.db.WithContext(ctx).Table("coupon_user as u").
		Select("`name`,`type`,`value`,min_price,start_at,end_at,`desc`,u.id").
		Scopes(model.OffsetPage(offset, limit), model.WhereRelease).
		Joins("left join coupon as c on u.coupon_id = c.id").
		Where("user_id=?", userID).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.coupon] by db uid: %v", userID)
	}
	return
}

//GetCouponByID 获取优惠券详情
func (r *Repo) GetCouponByID(ctx context.Context, id int) (*model.CouponModel, error) {
	coupon := new(model.CouponModel)
	err := r.db.WithContext(ctx).First(&coupon, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.coupon] first id: %v", id)
	}
	return coupon, nil
}
