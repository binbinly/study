package emoticon

import (
	"github.com/gin-gonic/gin"

	"chat-micro/app/logic/service"
	"chat-micro/pkg/app"
	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
)

// Cat 表情包所有分裂
// @Summary 表情包
// @Description 表情包
// @Tags 表情包
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @success 0 {object} app.Response{data=[]model.Emoticon} "调用成功结构"
// @Router /emoticon/list [get]
func Cat(c *gin.Context) {
	list, err := service.Svc.EmoticonCat(c.Request.Context())
	if err != nil {
		logger.Warnf("[http.emoticon] list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}
