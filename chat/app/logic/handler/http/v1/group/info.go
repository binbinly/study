package group

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

// Info 获取群信息
// @Summary 获取群信息
// @Description 获取群信息
// @Tags 群组
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param id query int true "群ID"
// @success 0 {object} app.Response{data=model.GroupInfo} "调用成功结构"
// @Router /group/info [get]
func Info(c *gin.Context) {
	gID := cast.ToUint32(c.Query("id"))
	if gID == 0 {
		app.Error(c, errno.ErrBind)
		return
	}
	info, err := service.Svc.GroupInfo(c.Request.Context(), app.GetUInt32UserID(c), gID)
	if errors.Is(err, service.ErrGroupNotFound) {
		app.Error(c, ecode.ErrGroupNotFound)
		return
	} else if errors.Is(err, service.ErrGroupUserNotJoin) {
		app.Error(c, ecode.ErrGroupNotJoin)
		return
	} else if err != nil {
		log.Warnf("[http.group] info err: %v", err)
		app.Error(c, ecode.ErrGroupNotFound)
		return
	}
	app.Success(c, info)
}
