package user

import (
	"errors"

	"github.com/gin-gonic/gin"

	"chat-micro/app/logic/ecode"
	"chat-micro/app/logic/service"
	"chat-micro/pkg/app"
	"chat-micro/pkg/logger"
)

// Profile 获取用户信息
// @Summary 获取个人资料
// @Description 获取个人资料
// @Tags 用户
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @success 0 {object} app.Response{data=model.User} "调用成功结构"
// @Router /user/profile [get]
func Profile(c *gin.Context) {
	user, err := service.Svc.UserInfoByID(c.Request.Context(), app.GetUInt32UserID(c))
	if errors.Is(err, service.ErrUserNotFound) {
		app.Error(c, ecode.ErrUserNotFound)
		return
	} else if errors.Is(err, service.ErrUserFrozen) {
		app.Error(c, ecode.ErrUserFrozen)
		return
	} else if err != nil {
		logger.Warnf("[http.user] profile get user info err: %v", err)
		app.Error(c, ecode.ErrUserNotFound)
		return
	}
	app.Success(c, user)
}
