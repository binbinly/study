package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
)

func HandleErrors(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("[gin.exception] err:%v", err)

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  errno.InternalServerError.Msg(),
			})
		}
	}()
	c.Next()
}
