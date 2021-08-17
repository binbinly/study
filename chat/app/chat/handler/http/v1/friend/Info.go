package friend

import (
	"chat/app/center"
	"chat/app/chat"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"chat/app/chat/ecode"
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
// @success 0 {object} app.Response{data=model.FriendInfo} "调用成功结构"
// @Router /friend/info [get]
func Info(c *gin.Context) {
	friendID := cast.ToUint32(c.Query("id"))
	if friendID == 0 {
		app.Error(c, errno.ErrBind)
		return
	}
	info, err := chat.Svc.FriendInfo(c.Request.Context(), app.GetUInt32UserID(c), friendID)
	if errors.Is(err, center.ErrUserNotFound) {
		app.Error(c, ecode.ErrUserNotFound)
		return
	} else if err != nil {
		log.Warnf("[http.friend] info err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, info)
}
