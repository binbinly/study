package moment

import (
	"chat/app/logic/handler/http"
	"errors"

	"github.com/gin-gonic/gin"

	"chat/app/logic/ecode"
	"chat/app/logic/service"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
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
		log.Warnf("[http.moment] timeline err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}
