package group

import (
	"errors"

	"github.com/gin-gonic/gin"

	"chat/app/logic/ecode"
	"chat/app/logic/service"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// Create 创建
// @Summary 创建群组
// @Description 创建群组
// @Tags 群组
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param ids body string true "用户id列表"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{}}"
// @Router /group/create [post]
func Create(c *gin.Context) {
	var req IdsParams
	valid := app.BindJson(c, &req)
	if !valid {
		app.Error(c, errno.ErrBind)
		return
	}
	err := service.Svc.GroupCreate(c, app.GetUInt32UserId(c), req.Ids)
	if errors.Is(err, service.ErrFriendNotRecord) {
		app.Error(c, ecode.ErrFriendNotFound)
		return
	} else if err != nil {
		log.Warnf("[http.group] create err: %v", err)
		app.Error(c, ecode.ErrGroupCreate)
		return
	}
	app.SuccessNil(c)
}
