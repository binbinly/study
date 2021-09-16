package middleware

import (
	"github.com/gin-gonic/gin"

	"mall/app/conf"
	"mall/pkg/app"
	"mall/pkg/errno"
	"mall/pkg/log"
)

// JWT 认证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		payload, err := app.ParseRequestToken(c, conf.Conf.App.JwtSecret)
		if err != nil {
			app.Error(c, errno.ErrInvalidToken)
			return
		}
		log.Debugf("context is: %+v", payload)

		// set uid to context
		c.Set("uid", payload.UserID)

		c.Next()
	}
}
