import request from '@/utils/request'

const api = {
  Login: '/login',
  Register: '/reg',
  Home: '/home',
  HomeSetting: '/home_setting',
  Notice: '/notice',
  HotKeyword: '/hot_keyword',
  Area: '/area',
  PayList: '/pay_list',
  Goods: {
    List: '/goods/list',
    Detail: '/goods/detail',
    Sku: '/goods/sku',
    Category: 'goods/category'
  },
  User: {
    Update: '/user/edit',
    Logout: '/user/logout'
  },
  UserAddress: {
    List: '/address/list',
    Add: '/address/add',
    Edit: '/address/edit',
    Del: '/address/del'
  },
  Cart: {
    List: '/cart/list',
    Del: '/cart/del',
    EditNum: '/cart/edit_num',
    Edit: '/cart/edit',
    Add: '/cart/add',
    Empty: '/cart/empty'
  },
  Coupon: {
    List: '/coupon/list',
    My: '/coupon/my',
    Draw: '/coupon/draw'
  },
  Order: {
    Submit: '/order/submit',
    GoodsSubmit: '/order/goods_submit',
    Detail: '/order/detail',
    List: '/order/list',
    Notify: '/order/notify',
    Refund: '/order/refund',
    Receipt: '/order/receipt',
    Cancel: '/order/cancel',
    Comment: '/order/comment'
  },
  post(url, data, params = null, auth = true, hideLoading = false) {
    return request({
      url,
      method: 'post',
      data,
      params,
      auth,
      hideLoading
    })
  },
  get(url, params = null, auth = false, hideLoading = true) {
    return request({
      url,
      params,
      auth,
      hideLoading
    })
  }
}

export default api
