<template>
  <div>
    <!-- 导航栏 -->
    <form>
      <van-search v-model="keyword" show-action :placeholder="placeholder" @search="onSearch" @cancel="onCancel" />
    </form>

    <div v-if="list.length === 0">
      <div class="py-2 flex align-center justify-center">
        <span class="text-light-muted">搜索指定内容</span>
      </div>

      <div class="px-4 flex flex-wrap">
        <div class="flex align-center justify-center mb-3 border-left border-right" style="width: 110px;" v-for="(item,index) in typeList"
             :key="index" @click="changeSearchType(item)">
          <span class="text-hover-primary">{{item.name}}</span>
        </div>
      </div>
    </div>
    <van-cell v-for="(item,index) in list" :title="item.name" center @click="openUserBase(item.id)" is-link>
      <template #icon>
        <van-image class="pr-1" round width="35" height="35" :src="item.avatar|formatAvatar" />
      </template>
    </van-cell>
  </div>
</template>

<script>
import { searchUser } from '@/api/common.js'
export default {
  data() {
    return {
      typeList: [{
        name: "聊天记录",
        key: "history"
      }, {
        name: "用户",
        key: "user"
      }, {
        name: "群聊",
        key: "group"
      }],
      searchType: "",
      keyword: "",
      list: []
    }
  },
  computed: {
    placeholder() {
      const obj = this.typeList.find((item) => {
        return item.key === this.searchType
      })
      if (obj) {
        return '搜索' + obj.name
      }
      return '请输入关键字'
    }
  },
  methods: {
    onSearch(val) {
      searchUser({
        keyword: this.keyword
      }).then(res => {
        this.list = res
      })
    },
    onCancel() {
      this.$router.back()
    },
    // 打开用户资料
    openUserBase(user_id) {
      this.$router.push({ path: '/user_base', query: { id: user_id } })
    },
    changeSearchType(item) {
      this.searchType = item.key
    }
  }
}
</script>

<style>
</style>
