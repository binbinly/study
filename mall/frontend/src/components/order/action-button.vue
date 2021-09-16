<template>
  <div>
    <template v-if="order.status == 1">
      <van-button type="warning" plain size="small" class="ml-1" @click.stop="openPay">去支付</van-button>
      <van-button type="warning" plain size="small" class="ml-1" @click.stop="onCancel">取消订单</van-button>
    </template>
    <template v-else-if="order.status === 2">
      <van-button type="warning" plain size="small" @click.stop="openRefund">申请退款</van-button>
    </template>
    <template v-else-if="order.status == 3">
      <van-button type="warning" plain size="small" class="ml-1" @click.stop="openLogistics">查看物流</van-button>
      <van-button type="warning" plain size="small" class="ml-1" @click.stop="onReceived">确认收货</van-button>
    </template>
    <template v-else-if="order.status == 4">
      <van-button type="warning" plain size="small" class="ml-1" @click.stop="openComment">去评价</van-button>
    </template>
  </div>
</template>

<script>
import { orderCancel, orderReceipt } from '@/api/order'
import { Dialog, Notify } from 'vant'
export default {
  props: {
    order: Object
  },
  methods: {
    openPay() {
      this.$router.push({ path: '/order_pay', query: { order_no: this.order.order_no, amount: this.order.amount } })
    },
    onCancel() {
      Dialog.confirm({
        title: '提示',
        message: '确定要取消吗？',
      })
        .then(() => {
          orderCancel({ order_no: this.order.order_no }).then(() => {
            this.$emit('cancel')
            Notify({ type: 'success', message: '取消成功' });
          }).catch(() => {
            Notify({ type: 'danger', message: '取消失败' });
          })
        })
        .catch(() => {
          // on cancel
        });
    },
    openRefund() {
      this.$router.push({ path: '/order_refund', query: { order_no: this.order.order_no } });
    },
    openLogistics() {
      this.$router.push({ path: '/order_logistics', query: { order_no: this.order.order_no } })
    },
    openComment() {
      const ids = this.order.goods.map(item => {
        return item.goods_id
      })
      this.$router.push({ path: '/order_comment', query: { order_no: this.order.order_no, ids: ids.join(',') } })
    },
    onReceived() {
      Dialog.confirm({
        title: '提示',
        message: '确定已经收到货了吗？',
      })
        .then(() => {
          orderReceipt({ order_no: this.order.order_no }).then(() => {
            Notify({ type: 'success', message: '取消成功' });
          }).catch(() => {
            Notify({ type: 'danger', message: '取消失败' });
          })
        })
        .catch(() => {
          // on cancel
        });
    }
  }
}
</script>

<style>
</style>
