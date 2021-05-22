package group

import (
	"chat/app/logic/service"
	"github.com/gin-gonic/gin"

	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// List 群组列表
// @Summary 群组列表
// @Description 群组列表
// @Tags 群组
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Success 200 {string} json "{"code":0,"message":"OK","data":null}"
// @Router /group/list [get]
func List(c *gin.Context) {
	list, err := service.Svc.GroupMyList(c, app.GetUInt32UserId(c))
	if err != nil {
		log.Warnf("[http.group] list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}
