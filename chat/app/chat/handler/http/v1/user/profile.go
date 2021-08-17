package user

import (
	"errors"

	"github.com/gin-gonic/gin"

	"chat/app/center"
	"chat/app/chat"
	"chat/app/chat/ecode"
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
// @success 0 {object} app.Response{data=model.UserInfo} "调用成功结构"
// @Router /user/profile [get]
func Profile(c *gin.Context) {
	user, err := chat.Svc.GetUserByID(c.Request.Context(), app.GetUInt32UserID(c))
	if errors.Is(err, center.ErrUserNotFound) {
		app.Error(c, ecode.ErrUserNotFound)
		return
	} else if errors.Is(err, center.ErrUserFrozen) {
		app.Error(c, ecode.ErrUserFrozen)
		return
	} else if err != nil {
		log.Warnf("[http.user] profile get user info err: %v", err)
		app.Error(c, ecode.ErrUserNotFound)
		return
	}
	app.Success(c, user)
}
