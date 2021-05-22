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

// User 获取群成员
// @Summary 获取群成员
// @Description 获取群成员
// @Tags 群组
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param id query int true "群ID"
// @Success 200 {string} json "{"code":0,"message":"OK","data":null}"
// @Router /group/user [get]
func User(c *gin.Context) {
	gId := cast.ToUint32(c.Query("id"))
	if gId == 0 {
		app.Error(c, errno.ErrBind)
		return
	}
	user, err := service.Svc.GroupUserAll(c, app.GetUInt32UserId(c), gId)
	if errors.Is(err, service.ErrGroupNotFound) {
		app.Error(c, ecode.ErrGroupNotFound)
		return
	} else if errors.Is(err, service.ErrGroupUserNotJoin) {
		app.Error(c, ecode.ErrGroupNotJoin)
		return
	} else if err != nil {
		log.Errorf("[http.group] user err: %v", err)
		app.Error(c, ecode.ErrGroupNotFound)
		return
	}
	app.Success(c, user)
}
