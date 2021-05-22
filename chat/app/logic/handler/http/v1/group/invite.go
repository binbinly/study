package group

import (
	"errors"
	"github.com/gin-gonic/gin"

	"chat/app/logic/ecode"
	"chat/app/logic/service"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// Invite 邀请好友
// @Summary 邀请好友
// @Description 邀请好友
// @Tags 群组
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param user body ActionParams true "The group info"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{}}"
// @Router /group/invite [post]
func Invite(c *gin.Context) {
	var req ActionParams
	valid := app.BindJson(c, &req)
	if !valid {
		app.Error(c, errno.ErrBind)
		return
	}
	err := service.Svc.GroupInviteUser(c, app.GetUInt32UserId(c), req.Id, req.UserId)
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
		log.Errorf("[http.group] invite err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}
