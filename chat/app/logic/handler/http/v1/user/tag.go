package user

import (
	"github.com/gin-gonic/gin"

	"chat/app/logic/service"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// Tag 标签列表
// @Summary 标签列表
// @Description 标签列表
// @Tags 用户
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Success 200 {string} json "{"code":0,"message":"OK","data":null}"
// @Router /user/tag [get]
func Tag(c *gin.Context) {
	list, err := service.Svc.UserTagList(c, app.GetUInt32UserId(c))
	if err != nil {
		log.Warnf("[http.user] tag err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}
