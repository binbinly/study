package user

import (
	"github.com/gin-gonic/gin"

	"chat/app/logic/ecode"
	"chat/app/logic/service"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// Profile 搜索用户
// @Summary 搜索用户
// @Description 搜索用户
// @Tags 用户
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param keyword body string true "搜索关键词"
// @Success 200 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /user/search [get]
func Search(c *gin.Context) {
	var req SearchParams
	v := app.BindJson(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	list, err := service.Svc.UserSearch(c, req.Keyword)
	if err != nil {
		log.Warnf("[http.user] search err: %v", err)
		app.Error(c, ecode.ErrUserNotFound)
		return
	}
	app.Success(c, list)
}
