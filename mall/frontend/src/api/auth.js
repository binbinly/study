import api from './index'

const authApi = {
  login(data) {
    return api.post(api.Login, data, null, false)
  },
  register(data) {
    return api.post(api.Register, data, null, false)
  }
}

export default authApi
