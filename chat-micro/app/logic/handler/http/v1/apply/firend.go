package apply

import (
	"errors"

	"github.com/gin-gonic/gin"

	"chat-micro/app/logic/ecode"
	"chat-micro/app/logic/service"
	"chat-micro/pkg/app"
	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
)

// Friend 申请好友
// @Summary 申请好友
// @Description 申请好友
// @Tags 好友申请
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body FriendParams true "friend"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /apply/friend [post]
func Friend(c *gin.Context) {
	var req FriendParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	userID := app.GetUInt32UserID(c)
	if userID == req.FriendID {
		app.Error(c, ecode.ErrUserNoSelf)
		return
	}
	err := service.Svc.ApplyFriend(c.Request.Context(), userID, req.FriendID, req.Nickname, req.LookMe, req.LookHim)
	if errors.Is(err, service.ErrApplyExisted) {
		app.Error(c, ecode.ErrApplyRepeatFailed)
		return
	} else if err != nil {
		logger.Warnf("[http.apply] friend err: %v", err)
		app.Error(c, ecode.ErrApplyFailed)
		return
	}
	app.SuccessNil(c)
}
