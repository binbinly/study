package idl

import "mall/app/model"

//TransferUserAddressList 转换用户售后地址输出
func TransferUserAddressList(list []*model.UserAddressModel) []*model.UserAddress {
	if len(list) == 0 {
		return []*model.UserAddress{}
	}
	res := make([]*model.UserAddress, 0, len(list))
	for _, addr := range list {
		res = append(res, &model.UserAddress{
			ID:        addr.ID,
			Name:      addr.Name,
			Phone:     addr.Phone,
			Province:  addr.Province,
			City:      addr.City,
			County:    addr.County,
			AreaCode:  addr.AreaCode,
			Detail:    addr.Detail,
			IsDefault: addr.IsDefault,
		})
	}
	return res
}
