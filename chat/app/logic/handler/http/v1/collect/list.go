package collect

import (
	"chat/app/logic/handler/http"
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
// @success 0 {object} app.Response{data=[]model.Collect} "调用成功结构"
// @Router /collect/list [get]
func List(c *gin.Context) {
	list, err := service.Svc.CollectGetList(c.Request.Context(), app.GetUInt32UserID(c), http.GetPageOffset(c))
	if err != nil {
		log.Warnf("[http.collect] list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}
