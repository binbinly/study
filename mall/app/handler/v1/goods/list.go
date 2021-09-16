package goods

import (
	"github.com/gin-gonic/gin"

	"mall/app/constvar"
	"mall/app/handler"
	"mall/app/service"
	"mall/pkg/app"
	"mall/pkg/errno"
	"mall/pkg/log"
)

// List 商品列表
// @Summary 商品列表
// @Description 商品列表
// @Tags 商品
// @Accept json
// @Produce json
// @Param p body int false "分页"
// @Param search body SearchParams true "search params"
// @success 0 {object} app.Response{data=[]model.GoodsList} "调用成功结构"
// @Router /goods/list [get]
func List(c *gin.Context) {
	var req SearchParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	list, err := service.Svc.GoodsList(c.Request.Context(), req.Cid, req.Cid, req.Keyword, req.Price, req.Field, req.Order, handler.GetPageOffset(c), constvar.DefaultLimit)
	if err != nil {
		log.Warnf("[v1.goods] list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}
