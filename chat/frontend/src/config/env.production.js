// 正式
module.exports = {
  title: 'chat',
  baseApi: process.env.VUE_APP_HTTP_URL,
  socketUrl: process.env.VUE_APP_WS_URL,
  uploadUrl: process.env.VUE_APP_DFS_URL
}
