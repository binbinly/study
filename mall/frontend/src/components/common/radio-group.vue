<template>
  <div class="row">
    <div class="span24-8 px-1 mb-1" :key="index" v-for="(item,index) in label.list">
      <div class="rounded px-1 py-1 bg-light-secondary font-sm text-center border"
           :class="label.selected === index ? 'radio-active':'border-light-secondary'" @click="changeRadio(index)">
        {{item.name}}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    label: Object,
  },
  methods: {
    changeRadio(index) {
      this.$emit('update:selected', index)
      this.$emit('change', index)
    }
  },
}
</script>

<style scoped>
.radio-active {
  background: #fce0d5 !important;
  color: #eb7320 !important;
  border-color: #eb7320 !important;
}
</style>
