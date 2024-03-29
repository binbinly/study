package collect

import (
	"chat/app/chat"
	"github.com/gin-gonic/gin"

	"chat/app/chat/ecode"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
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
	err := chat.Svc.CollectCreate(c.Request.Context(), req.Content, string(req.Options), app.GetUInt32UserID(c), req.Type)
	if err != nil {
		log.Warnf("[http.collect] create err: %v", err)
		app.Error(c, ecode.ErrCollectCreate)
		return
	}
	app.SuccessNil(c)
}
