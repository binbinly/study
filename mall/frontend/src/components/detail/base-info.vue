<template>
  <div class="p-1">
    <div class="font-lg">{{detail.title}}</div>
    <div class="font text-light-muted mb-2 line-h-sm">
      {{detail.intro}}
    </div>
    <price priceSize="font-lg" unitSize="font" :text="showPrice"></price>
  </div>
</template>

<script>
import price from "@/components/common/price.vue"
export default {
  components: {
    price
  },
  props: ['detail', 'showPrice'],
}
</script>

<style>
</style>
