package moment

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"chat-micro/app/logic/ecode"
	"chat-micro/app/logic/handler/http"
	"chat-micro/app/logic/service"
	"chat-micro/pkg/app"
	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
)

// List 动态列表
// @Summary 动态列表
// @Description 动态列表
// @Tags 朋友圈
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param user_id query int false "用户id"
// @Param p query int false "页码"
// @success 0 {object} app.Response{data=[]model.Moment} "调用成功结构"
// @Router /moment/list [get]
func List(c *gin.Context) {
	myID := app.GetUInt32UserID(c)
	userID := cast.ToUint32(c.Query("user_id"))
	if userID == 0 { // 默认查看自己的动态
		userID = myID
	}
	list, err := service.Svc.MomentList(c.Request.Context(), myID, userID, http.GetPageOffset(c))
	if errors.Is(err, service.ErrUserNotFound) {
		app.Error(c, ecode.ErrUserNotFound)
		return
	} else if err != nil {
		logger.Warnf("[http.moment] list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}
