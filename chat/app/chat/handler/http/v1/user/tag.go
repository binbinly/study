package user

import (
	"github.com/gin-gonic/gin"

	"chat/app/chat"
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
// @success 0 {object} app.Response{data=[]model.UserTag} "调用成功结构"
// @Router /user/tag [get]
func Tag(c *gin.Context) {
	list, err := chat.Svc.UserTagList(c.Request.Context(), app.GetUInt32UserID(c))
	if err != nil {
		log.Warnf("[http.user] tag err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}
