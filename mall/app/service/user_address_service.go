package service

import (
	"context"
	"github.com/pkg/errors"
	"mall/app/idl"
	"mall/app/model"
)

var (
	//ErrUserAddressNotFound 用户收货地址不存在
	ErrUserAddressNotFound = errors.New("user address not found")
)

//IUserAddress 收货地址服务接口
type IUserAddress interface {
	UserAddressAdd(ctx context.Context, userID int, name, phone,
		province, city, county, detail string, areaCode int, isDefault int8) (int, error)
	UserAddressEdit(ctx context.Context, id, userID int, m map[string]interface{}) error
	UserAddressList(ctx context.Context, userID int) ([]*model.UserAddress, error)
	DelUserAddress(ctx context.Context, id, userID int) error
}

//UserAddressAdd 添加收货地址
func (s *Service) UserAddressAdd(ctx context.Context, userID int, name, phone,
	province, city, county, detail string, areaCode int, isDefault int8) (int, error) {
	addr := &model.UserAddressModel{
		UID:       model.UID{UserID: userID},
		Name:      name,
		Phone:     phone,
		Province:  province,
		City:      city,
		County:    county,
		AreaCode:  areaCode,
		Detail:    detail,
		IsDefault: isDefault,
	}
	return s.repo.UserAddressCreate(ctx, addr)
}

//UserAddressEdit 收货地址修改
func (s *Service) UserAddressEdit(ctx context.Context, id, userID int, m map[string]interface{}) error {
	exist, err := s.repo.CheckUserAddress(ctx, id, userID)
	if err != nil {
		return errors.Wrapf(err, "[service.userAddress] check id : %v uid: %v", id, userID)
	}
	if !exist {
		return ErrUserAddressNotFound
	}
	err = s.repo.UserAddressUpdate(ctx, id, userID, m)
	if err != nil {
		return errors.Wrapf(err, "[service.userAddress] update id: %v, uid: %v", id, userID)
	}
	return nil
}

//UserAddressList 收货地址列表
func (s *Service) UserAddressList(ctx context.Context, userID int) ([]*model.UserAddress, error) {
	list, err := s.repo.GetUserAddressList(ctx, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.userAddress] list uid: %v", userID)
	}
	return idl.TransferUserAddressList(list), nil
}

//DelUserAddress 删除用户收货地址
func (s *Service) DelUserAddress(ctx context.Context, id, userID int) error {
	addr, err := s.repo.GetUserAddressByID(ctx, id, userID)
	if err != nil {
		return errors.Wrapf(err, "[service.userAddress] check id : %v uid: %v", id, userID)
	}
	if addr.ID == 0 {
		return ErrUserAddressNotFound
	}
	return s.repo.UserAddressDelete(ctx, addr)
}