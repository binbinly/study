package moment

import (
	"github.com/gin-gonic/gin"

	"chat-micro/app/logic/service"
	"chat-micro/pkg/app"
	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
)

// Create 发布
// @Summary 发布朋友圈
// @Description 发布朋友圈
// @Tags 朋友圈
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body CreateParams true "create"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /moment/create [post]
func Create(c *gin.Context) {
	var req CreateParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	err := service.Svc.MomentPush(c.Request.Context(), app.GetUInt32UserID(c), req.Content, req.Image, req.Video, req.Location, req.Type, req.SeeType, req.Remind, req.See)
	if err != nil {
		logger.Warnf("[http.moment] create err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}
