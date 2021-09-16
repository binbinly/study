<template>
  <div>
    <van-nav-bar :title="navTitle" left-arrow fixed placeholder @click-left="back" />

    <!-- 排序|筛选 -->
    <div class="d-flex border-bottom a-center position-fixed left-0 right-0 bg-white" style="height: 45px;z-index: 100;">
      <div class="flex-1 d-flex a-center j-center font-sm" v-for="(item,index) in screen.list" @click="changeScreen(index)">
        <span :class="screen.currentIndex === index ? 'main-text-color' : 'text-muted'">{{item.name}}</span>
        <div>
          <div class="iconfont icon-paixu-shengxu line-h0" :class="item.status==='asc'?'main-text-color':'text-light-muted'">
          </div>
          <div class="iconfont icon-paixu-jiangxu line-h0" :class="item.status==='desc'?'main-text-color':'text-light-muted'">
          </div>
        </div>
      </div>
      <div class="flex-1 d-flex a-center j-center main-text-color font-sm" @click="show = true">
        筛选
      </div>
    </div>

    <van-pull-refresh v-model="refreshing" @refresh="onRefresh" style="padding-top:60px;">
      <van-list v-model="loading" :finished="finished" finished-text="没有更多了" @load="onLoad" error-text="加载失败，请重试">
        <div class="row j-sb bg-white">
          <common-list v-for="(item,index) in list" :item="item" :index="index"></common-list>
        </div>
      </van-list>
    </van-pull-refresh>

    <van-popup v-model="show" position="right" :style="{ height: '100%' }">
      <div style="width:260px;">
        <card headTitle="价格范围" :headBorderBottom="false" :headTitleWeight="false">
          <!-- 单选按钮组 -->
          <zcm-radio-group :label="label" :selected.sync='label.selected'></zcm-radio-group>
        </card>
        <!-- 按钮 -->
        <div class="d-flex position-fixed bottom-0 right-0 w-100 border-top border-light-secondary">
          <van-button class="flex-1" color="#FD6801" type="primary" @click="confirm">确定</van-button>
          <van-button class="flex-1" type="default" @click="reset">重置</van-button>
        </div>
      </div>
    </van-popup>

  </div>
</template>

<script>
import zcmRadioGroup from "@/components/common/radio-group.vue"
import card from "@/components/common/card.vue"
import commonList from "@/components/common/common-list.vue"
import { goodsList } from '@/api/goods'
export default {
  components: {
    commonList,
    card,
    zcmRadioGroup
  },
  data() {
    return {
      show: false,
      navTitle: '商品列表',
      cat: 0,
      keyword: '',
      list: [],
      page: 1,
      loading: false,
      finished: false,
      refreshing: false,
      screen: {
        currentIndex: 0,
        list: [
          { name: "综合", status: 'desc', key: "sort" },
          { name: "销量", status: '', key: "sale_count" },
          { name: "价格", status: '', key: "price" },
        ]
      },
      label: {
        selected: 0,
        list: [
          { name: "不限", value: '' },
          { name: "0-50", value: "0,50" },
          { name: "50-100", value: "50,100" },
          { name: "100-500", value: "100,500" },
          { name: "500-1000", value: "500,1000" },
          { name: "大于1000", value: "1000" },
        ]
      },
    }
  },
  filters: {
    toString(value) {
      return value.toString();
    }
  },
  mounted() {
    this.cat = parseInt(this.$route.query.cat)
    this.navTitle = this.$route.query.title || '商品列表';
    if (this.$route.query.type === 'search') {
      this.keyword = this.navTitle
    }
  },
  methods: {
    back() {
      this.$router.back()
    },
    reset() {
      this.label.selected = 0
      this.show = false
    },
    confirm() {
      this.show = false
      this.onRefresh()
    },
    changeScreen(index) {
      // 判断当前点击是否已经是激活状态
      let oldIndex = this.screen.currentIndex
      let oldItem = this.screen.list[oldIndex]
      if (oldIndex === index) {
        oldItem.status = oldItem.status === 'asc' ? 'desc' : 'asc'
        // 加载数据
        this.onRefresh()
        return
      }
      let newIitem = this.screen.list[index]
      // 移除旧激活状态
      oldItem.status = ''
      this.screen.currentIndex = index
      // 增加新激活状态
      newIitem.status = 'desc'

      // 加载数据
      this.onRefresh()
    },
    onLoad() {
      this.getList()
    },
    onRefresh() {
      this.finished = false
      this.loading = true
      this.page = 1
      this.list = []
      this.onLoad()
    },
    getList() {
      const field = this.screen.list[this.screen.currentIndex].key
      const order = this.screen.list[this.screen.currentIndex].status
      const price = this.label.list[this.label.selected].value
      goodsList({ cat: this.cat, keyword: this.keyword, order, field, price }, this.page).then(res => {
        if (this.refreshing) {
          this.refreshing = false
        }
        this.page++
        this.loading = false
        this.list = [...this.list, ...res]
        if (res.length < 20) {
          this.finished = true
        }
      }).catch(err => {
        this.finished = true
      })
    },
  }
}
</script>

<style>
</style>
