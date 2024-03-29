package group

import (
	"chat/app/center"
	"chat/app/chat"
	"errors"

	"github.com/gin-gonic/gin"

	"chat/app/chat/ecode"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// KickOff 踢出群成员
// @Summary 踢出群成员
// @Description 踢出群成员
// @Tags 群组
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param user body ActionParams true "The group info"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /group/kickoff [post]
func KickOff(c *gin.Context) {
	var req ActionParams
	valid := app.BindJSON(c, &req)
	if !valid {
		app.Error(c, errno.ErrBind)
		return
	}
	err := chat.Svc.GroupKickOffUser(c.Request.Context(), app.GetUInt32UserID(c), req.ID, req.UserID)
	if errors.Is(err, chat.ErrGroupNotFound) {
		app.Error(c, ecode.ErrGroupNotFound)
		return
	} else if errors.Is(err, chat.ErrGroupUserNotJoin) {
		app.Error(c, ecode.ErrGroupNotJoin)
		return
	} else if errors.Is(err, center.ErrUserNotFound) {
		app.Error(c, ecode.ErrUserNotFound)
		return
	} else if errors.Is(err, chat.ErrGroupUserTargetNotJoin) {
		app.Error(c, ecode.ErrGroupSelectNotJoin)
		return
	} else if err != nil {
		log.Errorf("[http.group] kickoff err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}
