package user

import (
	"errors"

	"github.com/gin-gonic/gin"

	"chat-micro/app/logic/ecode"
	"chat-micro/app/logic/service"
	"chat-micro/pkg/app"
	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
)

// Report 好友举报
// @Summary 好友举报
// @Description 好友举报
// @Tags 好友
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body ReportParams true "report"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /report [post]
func Report(c *gin.Context) {
	var req ReportParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	userID := app.GetUInt32UserID(c)
	if userID == req.UserID {
		app.Error(c, ecode.ErrUserNoSelf)
		return
	}
	err := service.Svc.ReportCreate(c.Request.Context(), userID, req.UserID, req.Type, req.Category, req.Content)
	if errors.Is(err, service.ErrReportExisted) {
		app.Error(c, ecode.ErrReportHanding)
		return
	} else if err != nil {
		logger.Warnf("[http.user] report create err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}
