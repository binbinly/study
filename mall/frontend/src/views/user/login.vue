<template>
  <div>
    <div class="d-flex a-center j-sb py-2 px-3 text-light-muted">
      <div class="iconfont icon-shanchu1" @click="back"></div>
      <div class="font-md" @click="type = type == 'reg' ? 'login' : 'reg'">{{ type == 'reg' ? '去登录' : '去注册'}}</div>
    </div>

    <div class="pt-5">
      <div class="font-big m-1">{{ type == 'login' ? '密码登录' : '注册'}}</div>
      <van-form @submit="onSubmit">
        <template v-if="type == 'login'">
          <van-field v-model="username" name="username" label="用户名" placeholder="请输入用户名" :rules="[{ required: true }]" />
          <van-field v-model="password" type="password" name="password" label="密码" placeholder="请输入密码"
                     :rules="[{ required: true, validator: pwdValid, message: '密码长度6-20位' }]" />
        </template>
        <template v-else>
          <van-field v-model="username" name="username" label="用户名" placeholder="请输入用户名" :rules="[{ required: true }]" />
          <van-field v-model="phone" name="phone" label="手机号" placeholder="请输入手机号" :rules="[{ required: true }]" />
          <van-field v-model="password" type="password" name="password" label="密码" placeholder="请输入密码"
                     :rules="[{ required: true, validator: pwdValid, message: '密码长度6-20位' }]" />
          <van-field v-model="confirm_password" type="password" name="confirm_password" label="确认密码" placeholder="请确认密码"
                     :rules="[{ required: true, validator: confirmPwdValid, message: '密码长度6-20位' }]" />
        </template>
        <div style="margin: 16px;">
          <van-button class="text-white" :class="disabled ? 'main-bg-hover-color':'main-bg-color'" block native-type="submit">
            提交
          </van-button>
        </div>
        <label class="ml-2 checkbox d-flex a-center" @click="check = !check">
          <van-checkbox v-model="check" shape="square" />
          <span class="ml-1 font text-light-muted">已阅读并同意XXXXX协议</span>
        </label>
      </van-form>
    </div>

  </div>
</template>

<script>
import { mapMutations, mapActions } from 'vuex';
import authApi from '@/api/auth.js'
import { Toast } from 'vant'
export default {
  data() {
    return {
      type: "login",
      phone: "",
      username: "",
      password: "",
      confirm_password: "",
      check: false
    }
  },
  computed: {
    disabled() {
      if (this.username === '' || this.password === '') {
        return true
      }
      return false
    }
  },
  methods: {
    ...mapMutations(['login']),
    ...mapActions(['initCart']),
    back() {
      this.$router.back()
    },
    // 初始化表单
    initForm() {
      this.username = ''
      this.password = ''
      this.confirm_password = ''
      this.phone = ''
    },
    pwdValid(val) {
      return val.length >= 6 && val.length <= 20
    },
    confirmPwdValid(val) {
      return val === this.password
    },
    // 提交
    onSubmit() {
      if (this.type == 'login') { // 账号密码登录
        this.loginDo()
      } else { //注册
        this.register()
      }
    },
    loginDo() {
      authApi.login({
        username: this.username,
        password: this.password
      }).then(res => {
        // 修改vuex的state,持久化存储
        this.login(res)
        // 刷新购物车
        this.initCart()
        Toast('登录成功')
        this.back()
      })
    },
    register() {
      authApi.register({
        username: this.username,
        phone: this.phone,
        password: this.password,
        confirm_password: this.confirm_password,
      }).then(() => {
        Toast('注册成功')
        this.type = 'login'
        // 初始化表单
        this.initForm()
      })
    },
  }
}
</script>

<style>
</style>
