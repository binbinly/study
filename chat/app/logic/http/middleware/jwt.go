package middleware

import (
	"github.com/gin-gonic/gin"

	"chat/app/logic/conf"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
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
		log.Infof("context is: %+v", payload)

		// set uid to context
		c.Set("uid", payload.UserId)

		c.Next()
	}
}
