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

// Handle 处理好友申请
// @Summary 处理好友申请
// @Description 处理好友申请
// @Tags 好友申请
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param status body int true "状态 1=待处理 2=拒绝 3=同意 4=忽视"
// @Param nickname body string true "备注昵称"
// @Param look_me body int true "看我"
// @Param look_him body int true "看他"
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
		log.Warnf("[http.apply] handle err: %v", err)
		app.Error(c, ecode.ErrHandleFailed)
		return
	}
	app.SuccessNil(c)
}
