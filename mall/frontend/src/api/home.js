import api from './index'

//首页数据
export function homeData() {
  return api.get(api.Home)
}

//首页分类数据
export function homeCatData(cid = 0) {
  return api.get(api.HomeSetting, { cid })
}

//通知列表
export function noticeList(p) {
  return api.get(api.Notice, { p })
}

//搜索热词
export function hotSearch() {
  return api.get(api.HotKeyword)
}

//地区列表
export function areaList() {
  return api.get(api.Area)
}

//支付列表
export function getPayList() {
  return api.get(api.PayList)
}
