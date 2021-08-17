package emoticon

import (
	"chat/app/chat"
	"strings"

	"github.com/gin-gonic/gin"

	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// List 表情资源列表
// @Summary 表情包
// @Description 表情包
// @Tags 表情包
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param cat query string true "分类"
// @success 0 {object} app.Response{data=[]model.Emoticon} "调用成功结构"
// @Router /emoticon/list [get]
func List(c *gin.Context) {
	cat := strings.TrimSpace(c.Query("cat"))
	if cat == "" {
		app.Error(c, errno.ErrValidation)
		return
	}
	list, err := chat.Svc.Emoticon(c.Request.Context(), cat)
	if err != nil {
		log.Warnf("[http.emoticon] list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}
