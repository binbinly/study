package idl

import (
	pb "common/proto/member"
	"member/model"
)

//TransferAddressList 转换会员收货地址输出
func TransferAddressList(list []*model.MemberAddressModel) []*pb.Address {
	if len(list) == 0 {
		return []*pb.Address{}
	}
	res := make([]*pb.Address, 0, len(list))
	for _, addr := range list {
		res = append(res, TransferAddressItem(addr))
	}
	return res
}

//TransferAddressItem 转换会员收货地址输出
func TransferAddressItem(addr *model.MemberAddressModel) *pb.Address {
	return &pb.Address{
		Id:        addr.ID,
		Name:      addr.Name,
		Phone:     addr.Phone,
		Province:  addr.Province,
		City:      addr.City,
		County:    addr.County,
		AreaCode:  int64(addr.AreaCode),
		Detail:    addr.Detail,
		IsDefault: int32(addr.IsDefault),
	}
}
