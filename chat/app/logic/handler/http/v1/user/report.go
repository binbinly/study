package user

import (
	"errors"

	"github.com/gin-gonic/gin"

	"chat/app/logic/ecode"
	"chat/app/logic/service"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// Report 好友举报
// @Summary 好友举报
// @Description 好友举报
// @Tags 好友
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body ReportParams true "report"
// @Success 200 {string} json "{"code":0,"message":"OK","data":null}"
// @Router /report [post]
func Report(c *gin.Context) {
	var req ReportParams
	v := app.BindJson(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	userId := app.GetUInt32UserId(c)
	if userId == req.UserId {
		app.Error(c, ecode.ErrUserNoSelf)
		return
	}
	err := service.Svc.ReportCreate(c, userId, req.UserId, req.Type, req.Category, req.Content)
	if errors.Is(err, service.ErrReportExisted) {
		app.Error(c, ecode.ErrReportHanding)
		return
	} else if err != nil {
		log.Warnf("[http.user] report create err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}
