<template>
  <div>
    <div v-for="(item,index) in items" :key="index" @click.stop="openGoods(item.goods_id)">
      <div class="d-flex a-center pt-1 border-bottom bg-white border-light-secondary">
        <van-image class="pl-1 border border-light-secondary rounded flex-shrink" :src="item.cover" fit="cover" width="75" />

        <div class="flex-1 d-flex flex-column pl-1">
          <div class="text-dark font-sm">{{item.goods_name}}</div>
          <!-- 规格属性 -->
          <div class="d-flex text-light-muted" style="padding:5px 0;">
            {{item.sku_name || '默认'}}
          </div>
          <div class="mt-auto d-flex j-sb">
            <price :text="item.price" class="a-center" />
            <span class="mr-1 font-sm">x{{item.num}}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import price from "@/components/common/price.vue"
export default {
  components: {
    price
  },
  props: {
    items: Array
  },
  methods: {
    openGoods(id) {
      this.$router.push({ path: '/goods_detail', query: { id } })
    }
  }
}
</script>

<style>
</style>
