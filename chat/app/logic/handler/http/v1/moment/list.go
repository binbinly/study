package moment

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

// List 动态列表
// @Summary 动态列表
// @Description 动态列表
// @Tags 朋友圈
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param user_id query int false "用户id"
// @Param p query int false "页码"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{}"
// @Router /moment/list [get]
func List(c *gin.Context) {
	myId := app.GetUInt32UserId(c)
	userId := cast.ToUint32(c.Query("user_id"))
	if userId == 0 { // 默认查看自己的动态
		userId = myId
	}
	list, err := service.Svc.MomentList(c, myId, userId, app.GetPageOffset(c))
	if errors.Is(err, service.ErrUserNotFound) {
		app.Error(c, ecode.ErrUserNotFound)
		return
	} else if err != nil {
		log.Warnf("[http.moment] list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}
