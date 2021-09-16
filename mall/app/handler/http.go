package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"mall/app/constvar"
	"mall/pkg/utils"
)

//GetPageOffset 获取分页起始偏移量
func GetPageOffset(c *gin.Context) int {
	offset := 0
	page := cast.ToInt(c.Query("p"))
	if page > 0 {
		offset = (page - 1) * constvar.DefaultLimit
	}
	return offset
}

// healthCheckResponse 健康检查响应结构体
type healthCheckResponse struct {
	Status   string `json:"status"`
	Hostname string `json:"hostname"`
}

// HealthCheck will return OK if the underlying BoltDB is healthy. At least healthy enough for demoing purposes.
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, healthCheckResponse{Status: "UP", Hostname: utils.GetHostname()})
}
