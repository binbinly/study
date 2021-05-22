package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"chat/pkg/log"
)

// BindJson 绑定请求参数
func BindJson(c *gin.Context, form interface{}) bool {
	if err := c.ShouldBindJSON(form); err != nil {
		log.Infof("[bind.json] param err: %v", err)
		return false
	}
	return true
}

// GetUserId 返回用户id
func GetUserId(c *gin.Context) int {
	if c == nil {
		return 0
	}
	// uid 必须和 middleware/auth 中的 uid 命名一致
	return c.GetInt("uid")
}

func GetUInt32UserId(c *gin.Context) uint32 {
	return uint32(GetUserId(c))
}

// RouteNotFound 未找到相关路由
func RouteNotFound(c *gin.Context) {
	c.String(http.StatusNotFound, "the route not found")
}