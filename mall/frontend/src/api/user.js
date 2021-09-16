import api from './index'

//用户注销
export function userLogout() {
  return api.get(api.User.Logout, null, true)
}

//修改用户信息
export function userEdit(data) {
  return api.post(api.User.Edit, data)
}

//用户收货地址
export function userAddressList() {
  return api.get(api.UserAddress.List, null, true)
}

//添加用户收货地址
export function userAddressAdd(data) {
  return api.post(api.UserAddress.Add, data)
}

//修改用户收货地址
export function userAddressEdit(data) {
  return api.post(api.UserAddress.Edit, data)
}

//删除用户收货地址
export function userAddressDel(id) {
  return api.get(api.UserAddress.Del, { id }, true)
}
