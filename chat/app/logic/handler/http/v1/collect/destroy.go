package collect

import (
	"github.com/gin-gonic/gin"

	"chat/app/logic/ecode"
	"chat/app/logic/service"
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
// @Success 200 {string} json "{"code":0,"message":"OK","data":{}}"
// @Router /collect/destroy [post]
func Destroy(c *gin.Context) {
	var req DestroyParams
	v := app.BindJson(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	err := service.Svc.CollectDestroy(c, app.GetUInt32UserId(c), req.Id)
	if err != nil {
		log.Warnf("[http.collect] destroy err: %v", err)
		app.Error(c, ecode.ErrCollectDestroy)
		return
	}
	app.SuccessNil(c)
}
