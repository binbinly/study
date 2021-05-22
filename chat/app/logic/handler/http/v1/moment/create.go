package moment

import (
	"github.com/gin-gonic/gin"

	"chat/app/logic/service"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// Create 发布
// @Summary 发布朋友圈
// @Description 发布朋友圈
// @Tags 朋友圈
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body CreateParams true "create"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{}"
// @Router /moment/create [post]
func Create(c *gin.Context) {
	var req CreateParams
	v := app.BindJson(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	err := service.Svc.MomentPush(c, app.GetUInt32UserId(c), req.Content, req.Image, req.Video, req.Location, req.Type, req.SeeType, req.Remind, req.See)
	if err != nil {
		log.Warnf("[http.moment] create err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}
