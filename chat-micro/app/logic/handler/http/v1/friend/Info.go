package friend

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"chat-micro/app/logic/ecode"
	"chat-micro/app/logic/service"
	"chat-micro/pkg/app"
	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
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
	info, err := service.Svc.FriendInfo(c.Request.Context(), app.GetUInt32UserID(c), friendID)
	if errors.Is(err, service.ErrUserNotFound) {
		app.Error(c, ecode.ErrUserNotFound)
		return
	} else if err != nil {
		logger.Warnf("[http.friend] info err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, info)
}
