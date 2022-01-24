package middleware

import (
	"github.com/gin-gonic/gin"

	"chat-micro/app/logic/conf"
	"chat-micro/pkg/app"
	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
)

// JWT 认证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		payload, err := app.ParseRequestToken(c, conf.Conf.JwtSecret)
		if err != nil {
			app.Error(c, errno.ErrInvalidToken)
			return
		}
		logger.Debugf("context is: %+v", payload)

		// set uid to context
		c.Set("uid", payload.UserID)

		c.Next()
	}
}
