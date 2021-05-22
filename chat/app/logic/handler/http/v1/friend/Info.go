package friend

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"chat/app/logic/ecode"
	"chat/app/logic/service"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// Info 获取好友信息
// @Summary 获取好友信息
// @Description 获取好友信息
// @Tags 好友
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param id query int true "好友ID"
// @Success 200 {string} json "{"code":0,"message":"OK","data":null}"
// @Router /friend/info [get]
func Info(c *gin.Context) {
	friendId := cast.ToUint32(c.Query("id"))
	if friendId == 0 {
		app.Error(c, errno.ErrBind)
		return
	}
	info, err := service.Svc.FriendInfo(c, app.GetUInt32UserId(c), friendId)
	if errors.Is(err, service.ErrUserNotFound) {
		app.Error(c, ecode.ErrUserNotFound)
		return
	} else if err != nil {
		log.Warnf("[http.friend] info err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, info)
}
