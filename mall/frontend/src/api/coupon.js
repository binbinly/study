import api from './index'

//优惠券列表
export function getCouponList(p = 1) {
  return api.get(api.Coupon.List, { p }, true)
}

//我的优惠券
export function getMyCoupon(p = 1) {
  return api.get(api.Coupon.My, { p }, true)
}

//优惠券领取
export function couponDraw(id) {
  return api.get(api.Coupon.Draw, { id }, true)
}
