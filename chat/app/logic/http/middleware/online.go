package middleware

import (
	"chat/app/logic/ecode"
	"chat/app/logic/service"
	"github.com/gin-gonic/gin"

	"chat/pkg/app"
	"chat/pkg/log"
)

// Online 上线检查
func Online() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		log.Info("online start userId: ", app.GetUInt32UserId(c))
		is, err := service.Svc.CheckOnline(c, app.GetUInt32UserId(c))
		if err != nil {
			app.Error(c, ecode.ErrChatOffline)
			log.Warnf("online check userId:%v, err:%v", app.GetUInt32UserId(c), err)
			return
		}
		if !is {
			app.Error(c, ecode.ErrChatOffline)
			return
		}

		c.Next()
	}
}
