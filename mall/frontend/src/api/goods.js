import api from './index'

//商品列表
export function goodsList(data, p) {
  return api.post(api.Goods.List, data, { p }, false)
}

//商品详情
export function goodsDetail(id) {
  return api.get(api.Goods.Detail, { id })
}

//商品SKU
export function goodsDetailSku(id) {
  return api.get(api.Goods.Sku, { id })
}

//商品所有分类
export function categoryTree() {
  return api.get(api.Goods.Category)
}
