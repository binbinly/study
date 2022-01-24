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

// Handle 处理好友申请
// @Summary 处理好友申请
// @Description 处理好友申请
// @Tags 好友申请
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body HandleParams true "handle"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /apply/handle [post]
func Handle(c *gin.Context) {
	var req HandleParams
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
	err := service.Svc.ApplyHandle(c.Request.Context(), userID, req.FriendID, req.Nickname, req.LookMe, req.LookHim)
	if errors.Is(err, service.ErrApplyNotFound) {
		app.Error(c, ecode.ErrApplyNotFoundFailed)
		return
	} else if err != nil {
		logger.Warnf("[http.apply] handle err: %v", err)
		app.Error(c, ecode.ErrHandleFailed)
		return
	}
	app.SuccessNil(c)
}
