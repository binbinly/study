package moment

import (
	"errors"

	"github.com/gin-gonic/gin"

	"chat-micro/app/logic/ecode"
	"chat-micro/app/logic/handler/http"
	"chat-micro/app/logic/service"
	"chat-micro/pkg/app"
	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
)

// Timeline 我的朋友圈
// @Summary 我的朋友圈
// @Description 我的朋友圈
// @Tags 朋友圈
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param p query int false "页码"
// @success 0 {object} app.Response{data=[]model.Moment} "调用成功结构"
// @Router /moment/timeline [get]
func Timeline(c *gin.Context) {
	list, err := service.Svc.MomentTimeline(c.Request.Context(), app.GetUInt32UserID(c), http.GetPageOffset(c))
	if errors.Is(err, service.ErrUserNotFound) {
		app.Error(c, ecode.ErrUserNotFound)
		return
	} else if err != nil {
		logger.Warnf("[http.moment] timeline err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}
