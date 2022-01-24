package group

import (
	"errors"

	"github.com/gin-gonic/gin"

	"chat-micro/app/logic/ecode"
	"chat-micro/app/logic/service"
	"chat-micro/pkg/app"
	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
)

// Invite 邀请好友
// @Summary 邀请好友
// @Description 邀请好友
// @Tags 群组
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param user body ActionParams true "The group info"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /group/invite [post]
func Invite(c *gin.Context) {
	var req ActionParams
	valid := app.BindJSON(c, &req)
	if !valid {
		app.Error(c, errno.ErrBind)
		return
	}
	err := service.Svc.GroupInviteUser(c.Request.Context(), app.GetUInt32UserID(c), req.ID, req.UserID)
	if errors.Is(err, service.ErrGroupNotFound) {
		app.Error(c, ecode.ErrGroupNotFound)
		return
	} else if errors.Is(err, service.ErrGroupUserNotJoin) {
		app.Error(c, ecode.ErrGroupNotJoin)
		return
	} else if errors.Is(err, service.ErrUserNotFound) {
		app.Error(c, ecode.ErrUserNotFound)
		return
	} else if errors.Is(err, service.ErrGroupUserExisted) {
		app.Error(c, ecode.ErrGroupExisted)
		return
	} else if err != nil {
		logger.Errorf("[http.group] invite err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}
