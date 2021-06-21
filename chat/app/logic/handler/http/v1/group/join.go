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

// Join 加入群
// @Summary 加入群
// @Description 加入群
// @Tags 群组
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param id query int true "群ID"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /group/join [get]
func Join(c *gin.Context) {
	gID := cast.ToUint32(c.Query("id"))
	if gID == 0 {
		app.Error(c, errno.ErrBind)
		return
	}
	err := service.Svc.GroupJoin(c.Request.Context(), app.GetUInt32UserID(c), gID)
	if errors.Is(err, service.ErrGroupNotFound) {
		app.Error(c, ecode.ErrGroupNotFound)
		return
	} else if errors.Is(err, service.ErrGroupUserExisted) {
		app.Error(c, ecode.ErrGroupExisted)
		return
	} else if err != nil {
		log.Warnf("[http.group] join err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}
