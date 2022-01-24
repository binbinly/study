package group

import (
	"github.com/gin-gonic/gin"

	"chat-micro/app/logic/service"
	"chat-micro/pkg/app"
	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
)

// List 群组列表
// @Summary 群组列表
// @Description 群组列表
// @Tags 群组
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @success 0 {object} app.Response{data=[]model.GroupList} "调用成功结构"
// @Router /group/list [get]
func List(c *gin.Context) {
	list, err := service.Svc.GroupMyList(c.Request.Context(), app.GetUInt32UserID(c))
	if err != nil {
		logger.Warnf("[http.group] list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}
