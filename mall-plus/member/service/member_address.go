package service

import (
	"context"

	"github.com/pkg/errors"

	"common/errno"
	"common/orm"
	pb "common/proto/member"
	"member/idl"
	"member/model"
)

//IMemberAddress 收货地址服务接口
type IMemberAddress interface {
	MemberAddressAdd(ctx context.Context, MemberID int64, name, phone,
		province, city, county, detail string, areaCode int, isDefault int8) (int64, error)
	MemberAddressEdit(ctx context.Context, id, MemberID int64, name, phone,
		province, city, county, detail string, areaCode int, isDefault int8) error
	MemberAddressList(ctx context.Context, MemberID int64) ([]*pb.Address, error)
	MemberAddressDel(ctx context.Context, id, memberID int64) error
	GetMemberAddressInfo(ctx context.Context, id, memberID int64) (*pb.Address, error)
}

//MemberAddressAdd 添加收货地址
func (s *Service) MemberAddressAdd(ctx context.Context, MemberID int64, name, phone,
	province, city, county, detail string, areaCode int, isDefault int8) (int64, error) {
	addr := &model.MemberAddressModel{
		MID:       orm.MID{MemberID: MemberID},
		Name:      name,
		Phone:     phone,
		Province:  province,
		City:      city,
		County:    county,
		AreaCode:  areaCode,
		Detail:    detail,
		IsDefault: isDefault,
	}
	return s.repo.MemberAddressCreate(ctx, addr)
}

//MemberAddressEdit 收货地址修改
func (s *Service) MemberAddressEdit(ctx context.Context, id, MemberID int64, name, phone,
	province, city, county, detail string, areaCode int, isDefault int8) error {
	address, err := s.repo.GetMemberAddressByID(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[service.memberAddress] get id: %v", id)
	}
	if address == nil || address.MemberID != MemberID {
		return errno.ErrMemberAddressNotFound
	}
	isEdit := false
	if name != "" && address.Name != name {
		address.Name = name
		isEdit = true
	}
	if phone != "" && address.Phone != phone {
		address.Phone = phone
		isEdit = true
	}
	if province != "" && address.Province != province {
		address.Province = province
		isEdit = true
	}
	if city != "" && address.City != city {
		address.City = city
		isEdit = true
	}
	if county != "" && address.County != county {
		address.County = city
		isEdit = true
	}
	if detail != "" && address.Detail != detail {
		address.Detail = detail
		isEdit = true
	}
	if areaCode != 0 && address.AreaCode != areaCode {
		address.AreaCode = areaCode
		isEdit = true
	}
	if address.IsDefault != isDefault {
		address.IsDefault = isDefault
		isEdit = true
	}
	if isEdit {
		if err = s.repo.MemberAddressSave(ctx, address); err != nil {
			return errors.Wrapf(err, "[service.memberAddress] update id: %v, uid: %v", id, MemberID)
		}
	}

	return nil
}

//MemberAddressList 收货地址列表
func (s *Service) MemberAddressList(ctx context.Context, MemberID int64) ([]*pb.Address, error) {
	list, err := s.repo.GetMemberAddressList(ctx, MemberID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.memberAddress] list uid: %v", MemberID)
	}
	return idl.TransferAddressList(list), nil
}

//MemberAddressDel 删除用户收货地址
func (s *Service) MemberAddressDel(ctx context.Context, id, memberID int64) error {
	addr, err := s.repo.GetMemberAddressByID(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[service.memberAddress] check id : %v uid: %v", id, memberID)
	}
	if addr == nil || addr.MemberID != memberID {
		return errno.ErrMemberAddressNotFound
	}
	return s.repo.MemberAddressDelete(ctx, addr)
}

//GetMemberAddressInfo 获取收货地址信息
func (s *Service) GetMemberAddressInfo(ctx context.Context, id, memberID int64) (*pb.Address, error) {
	addr, err := s.repo.GetMemberAddressByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.memberAddress] check id : %v uid: %v", id, memberID)
	}
	if addr == nil || addr.MemberID != memberID {
		return nil, errno.ErrMemberAddressNotFound
	}
	return idl.TransferAddressItem(addr), nil
}
