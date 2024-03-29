package user

import (
	"github.com/gin-gonic/gin"

	"chat-micro/app/logic/ecode"
	"chat-micro/app/logic/service"
	"chat-micro/pkg/app"
	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
)

// Search 搜索用户
// @Summary 搜索用户
// @Description 搜索用户
// @Tags 用户
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param keyword body string true "搜索关键词"
// @success 0 {object} app.Response{data=[]model.User} "调用成功结构"
// @Router /user/search [get]
func Search(c *gin.Context) {
	var req SearchParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	list, err := service.Svc.UserSearch(c.Request.Context(), req.Keyword)
	if err != nil {
		logger.Warnf("[http.user] search err: %v", err)
		app.Error(c, ecode.ErrUserNotFound)
		return
	}
	app.Success(c, list)
}
