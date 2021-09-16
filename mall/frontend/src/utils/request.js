import axios from 'axios'
import $store from '@/store'
import $router from '@/router/index.js'
import { Toast } from 'vant'
// 根据环境不同引入不同api地址
import { baseApi } from '@/config'

// create an axios instance
const service = axios.create({
  baseURL: baseApi, // url = base api url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 5000, // request timeout
  headers: { 'Content-Type': 'application/json;charset=UTF-8' }
})

// request拦截器 request interceptor
service.interceptors.request.use(
  config => {
    // 不传递默认开启loading
    if (!config.hideLoading) {
      // loading
      Toast.loading({
        forbidClick: true
      })
    }
    if (config.auth === true) {
      const token = $store.getters.token
      if (!token) {
        return $router.replace({ path: '/login' })
      }
      config.headers['Token'] = token || ''
    }
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)
// response拦截器
service.interceptors.response.use(
  response => {
    Toast.clear()
    if (response.status === 200) {
      const result = response.data
      if (result.code === 0) {
        return Promise.resolve(result.data)
      } else if (result.code == 10108) {
        Toast('令牌已过期，请重新登录')
        $store.commit('logout')
        $router.push({ path: '/login' })
        return Promise.reject(result.msg)
      }
      Toast.fail(result.msg)
      return Promise.reject(result.msg)
    } else {
      Toast.fail('网络开小差了')
      console.log('response err', response)
      return Promise.reject(response.statusText)
    }
  },
  error => {
    Toast.clear()
    if (error.message === 'Network Error') {
      Toast.fail('服务器连接异常，请检查服务器！')
      return
    }
    console.log('err' + error) // for debug

    Toast.fail(error.message)
    return Promise.reject(error)
  }
)

export default service
