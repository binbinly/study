package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mall/pkg/errno"
)

// Response api的返回结构体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 成功返回
func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code: errno.Success.Code(),
		Msg:  errno.Success.Msg(),
		Data: data,
	})
}

// SuccessNil 成功返回，无数据
func SuccessNil(c *gin.Context) {
	Success(c, nil)
}

// Error 错误返回
func Error(c *gin.Context, err *errno.Error) {
	code, msg := errno.DecodeErr(err)
	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: gin.H{},
	})
}
