package repo

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"mall/app/model"
)

//IUserAddress 用户收货地址接口
type IUserAddress interface {
	UserAddressCreate(ctx context.Context, addr *model.UserAddressModel) (int, error)
	UserAddressUpdate(ctx context.Context, id, userID int, m map[string]interface{}) error
	GetUserAddressList(ctx context.Context, userID int) (list []*model.UserAddressModel, err error)
	GetUserAddressByID(ctx context.Context, id, userID int) (*model.UserAddressModel, error)
	CheckUserAddress(ctx context.Context, id, userID int) (bool, error)
	UserAddressDelete(ctx context.Context, addr *model.UserAddressModel) error
}

//UserAddressCreate 创建收货地址
func (r *Repo) UserAddressCreate(ctx context.Context, addr *model.UserAddressModel) (int, error) {
	if addr.IsDefault == 1 {
		if err := r.cancelDefaultAddress(ctx, addr.UserID); err != nil {
			return 0, err
		}
	}
	err := r.db.WithContext(ctx).Create(addr).Error
	if err != nil {
		return 0, errors.Wrapf(err, "[repo.userAddress] create")
	}
	return addr.ID, nil
}

// UserAddressUpdate 更新用户收货地址
func (r *Repo) UserAddressUpdate(ctx context.Context, id, userID int, m map[string]interface{}) error {
	if def, ok := m["is_default"]; ok {
		defInt := def.(int8)
		if defInt == 1 { //设置当前为默认收货地址，取消其他默认
			if err := r.cancelDefaultAddress(ctx, userID); err != nil {
				return err
			}
		}
	}
	err := r.db.WithContext(ctx).Model(&model.UserAddressModel{}).Where("id=? and user_id=?", id, userID).Updates(m).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.userAddress] update")
	}
	return nil
}

//GetUserAddressList 用户收货地址列表
func (r *Repo) GetUserAddressList(ctx context.Context, userID int) (list []*model.UserAddressModel, err error) {
	err = r.db.WithContext(ctx).Where("user_id=?", userID).Order(model.DefaultOrder).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.userAddress] find")
	}
	return
}

//GetUserAddressByID 用户收货地址详情
func (r *Repo) GetUserAddressByID(ctx context.Context, id, userID int) (*model.UserAddressModel, error) {
	addr := new(model.UserAddressModel)
	err := r.db.WithContext(ctx).Where("id=? and user_id=?", id, userID).First(addr).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.userAddress] first")
	}
	return addr, nil
}

//CheckUserAddress 检查用户收货地址是否存在
func (r *Repo) CheckUserAddress(ctx context.Context, id, userID int) (bool, error) {
	var c int64
	err := r.db.WithContext(ctx).Model(&model.UserAddressModel{}).Where("id=? and user_id=?", id, userID).Count(&c).Error
	if err != nil {
		return false, errors.Wrapf(err, "[repo.userAddress] count")
	}
	return c > 0, nil
}

//UserAddressDelete 删除收货地址
func (r *Repo) UserAddressDelete(ctx context.Context, addr *model.UserAddressModel) error {
	err := r.db.WithContext(ctx).Delete(addr).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.userAddress] delete")
	}

	return nil
}

//cancelDefaultAddress 取消默认地址
func (r *Repo) cancelDefaultAddress(ctx context.Context, userID int) error {
	err := r.db.WithContext(ctx).Model(&model.UserAddressModel{}).
		Where("user_id=? and is_default=1", userID).Update("is_default", 0).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.userAddress] set other not defualt by uid: %v", userID)
	}
	return nil
}