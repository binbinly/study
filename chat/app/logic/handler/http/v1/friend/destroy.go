package friend

import (
	"errors"

	"github.com/gin-gonic/gin"

	"chat/app/logic/ecode"
	"chat/app/logic/service"
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
// @Success 200 {string} json "{"code":0,"message":"OK","data":null}"
// @Router /friend/auth [post]
func Destroy(c *gin.Context) {
	var req DestroyParams
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
	err := service.Svc.FriendDestroy(c, userId, req.UserId)
	if errors.Is(err, service.ErrFriendNotRecord) {
		app.Error(c, ecode.ErrFriendNotFound)
		return
	} else if err != nil {
		log.Warnf("[http.friend] destroy err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}
