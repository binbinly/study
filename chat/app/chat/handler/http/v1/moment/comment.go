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

// Comment 评论
// @Summary 评论
// @Description 评论
// @Tags 朋友圈
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body CommentParams true "create"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /moment/comment [post]
func Comment(c *gin.Context) {
	var req CommentParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	err := chat.Svc.MomentComment(c.Request.Context(), app.GetUInt32UserID(c), req.ReplyID, req.ID, req.Content)
	if errors.Is(err, chat.ErrMomentNotFound) {
		app.Error(c, ecode.ErrMomentNotFound)
		return
	} else if errors.Is(err, center.ErrUserNotFound) {
		app.Error(c, ecode.ErrUserNotFound)
		return
	} else if err != nil {
		log.Warnf("[http.moment] comment err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}
