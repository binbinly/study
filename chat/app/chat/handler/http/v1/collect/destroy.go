package collect

import (
	"chat/app/chat"
	"github.com/gin-gonic/gin"

	"chat/app/chat/ecode"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// Destroy 删除收藏
// @Summary 删除收藏
// @Description 删除收藏
// @Tags 用户收藏
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body DestroyParams true "destroy"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /collect/destroy [post]
func Destroy(c *gin.Context) {
	var req DestroyParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	err := chat.Svc.CollectDestroy(c.Request.Context(), app.GetUInt32UserID(c), req.ID)
	if err != nil {
		log.Warnf("[http.collect] destroy err: %v", err)
		app.Error(c, ecode.ErrCollectDestroy)
		return
	}
	app.SuccessNil(c)
}
