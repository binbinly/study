<template>
  <div class="d-flex">
    <van-image width="186" height="265" :src="resdata[0]|formatImage" fit="cover" lazy-load @click="event" />
    <div class="d-flex flex-column">
      <van-image width="186" height="132" :src="resdata[1]|formatImage" fit="cover" lazy-load @click="event" />
      <van-image width="186" height="132" :src="resdata[2]|formatImage" fit="cover" lazy-load @click="event" />
    </div>
  </div>
</template>

<script>
import { Toast } from 'vant'
export default {
  props: {
    resdata: [Array, Object],
  },
  methods: {
    event() {
      Toast('广告点击')
    }
  }
}
</script>

<style>
</style>
