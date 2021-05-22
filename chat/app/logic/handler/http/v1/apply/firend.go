package apply

import (
	"errors"

	"github.com/gin-gonic/gin"

	"chat/app/logic/ecode"
	"chat/app/logic/service"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// Friend 申请好友
// @Summary 申请好友
// @Description 申请好友
// @Tags 好友申请
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param friend_id body int true "好友ID"
// @Param nickname body string true "备注昵称"
// @Param look_me body int true "看我"
// @Param look_him body int true "看他"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6Ik"}}"
// @Router /apply/friend [post]
func Friend(c *gin.Context) {
	var req FriendParams
	v := app.BindJson(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	userId := app.GetUInt32UserId(c)
	if userId == req.FriendId {
		app.Error(c, ecode.ErrUserNoSelf)
		return
	}
	err := service.Svc.ApplyFriend(c, userId, req.FriendId, req.Nickname, req.LookMe, req.LookHim)
	if errors.Is(err, service.ErrApplyExisted) {
		app.Error(c, ecode.ErrApplyRepeatFailed)
		return
	} else if err != nil {
		log.Warnf("[http.apply] friend err: %v", err)
		app.Error(c, ecode.ErrApplyFailed)
		return
	}
	app.SuccessNil(c)
}
