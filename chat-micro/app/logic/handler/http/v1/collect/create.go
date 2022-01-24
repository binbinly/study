package collect

import (
	"github.com/gin-gonic/gin"

	"chat-micro/app/logic/ecode"
	"chat-micro/app/logic/service"
	"chat-micro/pkg/app"
	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
)

// Create 添加收藏
// @Summary 添加收藏
// @Description 添加收藏
// @Tags 用户收藏
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param type body int true "聊天信息类型"
// @Param content body string true "内容"
// @Param options body []byte false "额外选项"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /collect/create [post]
func Create(c *gin.Context) {
	var req CreateParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	err := service.Svc.CollectCreate(c.Request.Context(), req.Content, string(req.Options), app.GetUInt32UserID(c), req.Type)
	if err != nil {
		logger.Warnf("[http.collect] create err: %v", err)
		app.Error(c, ecode.ErrCollectCreate)
		return
	}
	app.SuccessNil(c)
}
