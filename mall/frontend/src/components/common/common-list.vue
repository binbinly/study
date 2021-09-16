<template>
  <div class="px-1" style="width:44%" @click="openDetail">
    <van-image style="display:flex" :src="item.cover" fit="cover" />
    <div class="p-1 pt-1">
      <div><span class="text-inline font-sm">{{item.title}}</span></div>
      <span class="d-block font-small text-light-muted text-inline-3">{{item.intro}}</span>
      <div class="d-flex py-1">
        <price :text="item.price"></price>
        <div class="ml-1 a-self-end line-h"><span class="font-sm text-light-muted line-through">￥{{item.original_price}}</span></div>
      </div>
    </div>
  </div>
</template>

<script>
import price from "@/components/common/price.vue";
export default {
  components: {
    price
  },
  props: {
    item: Object,
    index: [Number, String]
  },
  methods: {
    openDetail() {
      this.$router.push({ path: '/goods_detail', query: { id: this.item.id } })
    }
  }
}
</script>

<style>
</style>
