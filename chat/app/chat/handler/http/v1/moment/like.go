package moment

import (
	"chat/app/center"
	"chat/app/chat"
	"errors"

	"github.com/gin-gonic/gin"

	"chat/app/chat/ecode"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// Like 点赞
// @Summary 点赞
// @Description 点赞
// @Tags 朋友圈
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body LikeParams true "create"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /moment/like [post]
func Like(c *gin.Context) {
	var req LikeParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	err := chat.Svc.MomentLike(c.Request.Context(), app.GetUInt32UserID(c), req.ID)
	if errors.Is(err, chat.ErrMomentNotFound) {
		app.Error(c, ecode.ErrMomentNotFound)
		return
	} else if errors.Is(err, center.ErrUserNotFound) {
		app.Error(c, ecode.ErrUserNotFound)
		return
	} else if err != nil {
		log.Warnf("[http.moment] like err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}
