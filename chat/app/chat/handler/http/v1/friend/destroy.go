package friend

import (
	"chat/app/chat"
	"errors"

	"github.com/gin-gonic/gin"

	"chat/app/chat/ecode"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// Destroy 删除好友
// @Summary 删除好友
// @Description 删除好友
// @Tags 好友
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body DestroyParams true "destroy"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /friend/auth [post]
func Destroy(c *gin.Context) {
	var req DestroyParams
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
	err := chat.Svc.FriendDestroy(c.Request.Context(), userID, req.UserID)
	if errors.Is(err, chat.ErrFriendNotRecord) {
		app.Error(c, ecode.ErrFriendNotFound)
		return
	} else if err != nil {
		log.Warnf("[http.friend] destroy err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}
