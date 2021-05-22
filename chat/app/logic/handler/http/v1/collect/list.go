package collect

import (
	"github.com/gin-gonic/gin"

	"chat/app/logic/service"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// List 收藏列表
// @Summary 收藏列表
// @Description 收藏列表
// @Tags 用户收藏
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param p query int false "页码"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{}"
// @Router /collect/list [get]
func List(c *gin.Context) {
	list, err := service.Svc.CollectGetList(c, app.GetUInt32UserId(c), app.GetPageOffset(c))
	if err != nil {
		log.Warnf("[http.collect] list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}
