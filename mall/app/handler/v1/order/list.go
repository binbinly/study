package order

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"mall/app/constvar"
	"mall/app/handler"
	"mall/app/service"
	"mall/pkg/app"
	"mall/pkg/errno"
	"mall/pkg/log"
)

// List 订单列表
// @Summary 订单列表
// @Description 订单列表
// @Tags 订单
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param p body int false "分页"
// @success 0 {object} app.Response{data=[]model.GoodsList} "调用成功结构"
// @Router /order/list [get]
func List(c *gin.Context) {
	status := cast.ToInt(c.Query("status"))
	list, err := service.Svc.MyOrderList(c.Request.Context(), app.GetUserID(c), status, handler.GetPageOffset(c), constvar.DefaultLimit)
	if err != nil {
		log.Warnf("[v1.order] list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}