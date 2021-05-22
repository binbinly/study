import api from './index'
import upload from '@/utils/upload'

export function login(data) {
  return api.post(api.Login, data, false)
}

export function reg(data) {
  return api.post(api.Reg, data, false)
}

export function searchUser(data) {
  return api.post(api.SearchUser, data)
}

export function httpGet(url) {
  return api.get(url)
}

export function uploadFile(file) {
  let data = new FormData()
  data.append('file', file)
  return upload({
    method: 'post',
    data
  })
}
