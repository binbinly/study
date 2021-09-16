package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mall/pkg/errno"
	"mall/pkg/log"
)

func HandleErrors(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("[gin.exception] err:%v", err)

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  errno.InternalServerError.Msg(),
			})
		}
	}()
	c.Next()
}
