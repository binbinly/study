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

// Quit 删除并退出群
// @Summary 删除并退出群
// @Description 删除并退出群
// @Tags 群组
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param id query int true "群ID"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /group/quit [get]
func Quit(c *gin.Context) {
	gID := cast.ToUint32(c.Query("id"))
	if gID == 0 {
		app.Error(c, errno.ErrBind)
		return
	}
	err := service.Svc.GroupUserQuit(c.Request.Context(), app.GetUInt32UserID(c), gID)
	if errors.Is(err, service.ErrGroupNotFound) {
		app.Error(c, ecode.ErrGroupNotFound)
		return
	} else if errors.Is(err, service.ErrGroupUserNotJoin) {
		app.Error(c, ecode.ErrGroupNotJoin)
		return
	} else if err != nil {
		log.Warnf("[http.group] quit err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}
