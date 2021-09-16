import { Toast, Dialog } from 'vant'
import { updateCartBadge } from '@/utils/index.js'
import * as api from '@/api/cart'
export default {
  state: {
    list: []
  },
  mutations: {
    // 初始化list
    initCartList(state, list) {
      state.list = list
      updateCartBadge(state.list.length)
    },
    // 删除选中
    delGoods(state, index) {
      state.list.splice(index, 1)
      updateCartBadge(state.list.length)
    },
    delBatchGoods(state, indexs) {
      //逆向循环
      for (let i = indexs.length - 1; i >= 0; i--) {
        state.list.splice(i, 1)
      }
      updateCartBadge(state.list.length)
    },
    // 加入购物车
    addGoodsToCart(state, goods) {
      for (const key in state.list) {
        if (state.list[key].goods_id == goods.goods_id && state.list[key].sku_id == goods.sku_id) {
          state.list[key].num = goods.num
          return
        }
      }
      state.list.unshift(goods)
      updateCartBadge(state.list.length)
    },
    // 清空购物车
    clearCart(state) {
      state.list = []
      updateCartBadge(0)
    }
  },
  actions: {
    // 更新购物车列表
    initCart({ commit }) {
      api.getCartList().then(result => {
        console.log('cart', result)
        // 赋值
        commit('initCartList', result)
      })
    },
    doAddCart({ commit }, { goods_id, sku_id, num }) {
      api.cartAdd({ goods_id, sku_id, num }).then(res => {
        commit('addGoodsToCart', res)
        Toast.success('添加成功')
      })
    },
    doDelCart({ state, commit }, index) {
      const cart = state.list[index]
      if (!cart) {
        return Toast('已删除哦')
      }
      Dialog.confirm({
        title: '提示',
        message: '确定要移除该商品吗？'
      })
        .then(() => {
          api.cartDelete({ id: cart.id }).then(() => {
            commit('delGoods', index)
            Toast.success('删除成功')
          })
        })
        .catch(() => {
          // on cancel
        })
    },
    doEmptyCart({ commit }) {
      Dialog.confirm({
        title: '提示',
        message: '确定清空购物车吗？'
      })
        .then(() => {
          api.cartEmpty().then(() => {
            commit('clearCart')
            Toast.success('清空成功')
          })
        })
        .catch(() => {
          // on cancel
        })
    },
    doEditCart({ state }, { index, item, sku_id, num }) {
      if (!sku_id) {
        return api.cartEditNum({ num, id: item.id }).then(() => {
          state.list[index].num = num
        })
      }
      api.cartEdit({ sku_id, num, id: item.id }).then(res => {
        state.list[index].sku_name = res.sku_name
        state.list[index].sku_id = res.sku_id
        state.list[index].num = res.num
      })
    }
  }
}
