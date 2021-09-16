<template>
  <div>
    <div @click="show=true">
      <div class="scroll-row-item" style="width: 24%;height: 55px;" v-for="(value,name) in attrs">
        <div class="d-flex flex-column a-center j-center">
          <div class="iconfont line-h-sm icon-cpu"></div>
          <div class="line-h-sm">{{name}}</div>
          <div class="text-light-muted line-h-sm">{{value}}</div>
        </div>
      </div>
    </div>

    <van-popup v-model="show" position="bottom" round>
      <div class="d-flex a-center j-center font-md border-bottom border-light-secondary" style="height: 50px;">
        参数配置
      </div>
      <div style="height: 400px;overflow: auto;">
        <van-cell :title="name" :value="value" v-for="(value,name) in baseAttrs" />
      </div>
      <van-button class="w-100" type="default" color="#FD6801" @click="show = false">确定</van-button>
    </van-popup>
  </div>
</template>

<script>
export default {
  props: ['baseAttrs'],
  data() {
    return {
      show: false
    }
  },
  computed: {
    attrs() {
      let i = 0;
      let attr = {}
      for (const key in this.baseAttrs) {
        if (i > 3) continue;
        attr[key.substring(0, 8)] = this.baseAttrs[key].substring(0, 8)
        i++
      }
      return attr
    }
  }
}
</script>

<style>
</style>
