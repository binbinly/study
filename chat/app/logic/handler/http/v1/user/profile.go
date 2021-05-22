package user

import (
	"github.com/gin-gonic/gin"

	"chat/app/logic/ecode"
	"chat/app/logic/service"
	"chat/pkg/app"
	"chat/pkg/log"
)

// Profile 获取用户信息
// @Summary 获取个人资料
// @Description 获取个人资料
// @Tags 用户
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Success 200 {string} json "{"code":0,"message":"OK","data":null}"
// @Router /user/profile [get]
func Profile(c *gin.Context) {
	user, err := service.Svc.UserInfoById(c, app.GetUInt32UserId(c))
	if err != nil {
		log.Warnf("[http.user] profile get user info err: %v", err)
		app.Error(c, ecode.ErrUserNotFound)
		return
	}
	app.Success(c, user)
}
