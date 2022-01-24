package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"chat-micro/pkg/logger"
)

// BindJSON 绑定请求参数
func BindJSON(c *gin.Context, form interface{}) bool {
	if err := c.ShouldBindJSON(form); err != nil {
		logger.Infof("[bind.json] param err: %v", err)
		return false
	}
	return true
}

// GetUserID 返回用户id
func GetUserID(c *gin.Context) int {
	if c == nil {
		return 0
	}
	// uid 必须和 middleware/auth 中的 uid 命名一致
	return c.GetInt("uid")
}

//GetUInt32UserID 获取uint32用户id
func GetUInt32UserID(c *gin.Context) uint32 {
	return uint32(GetUserID(c))
}

// RouteNotFound 未找到相关路由
func RouteNotFound(c *gin.Context) {
	c.String(http.StatusNotFound, "the route not found")
}