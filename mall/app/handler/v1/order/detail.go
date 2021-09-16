package order

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"mall/app/ecode"
	"mall/app/service"
	"mall/pkg/app"
	"mall/pkg/errno"
	"mall/pkg/log"
)

// Detail 订单详情
// @Summary 订单详情
// @Description 订单详情
// @Tags 订单
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param id body int true "订单id"
// @success 0 {object} app.Response{data=[]model.Order} "调用成功结构"
// @Router /order/detail [get]
func Detail(c *gin.Context) {
	id := cast.ToInt(c.Query("id"))
	if id == 0 {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	info, err := service.Svc.OrderDetail(c.Request.Context(), id, app.GetUserID(c))
	if errors.Is(err, service.ErrOrderNotFound) {
		app.Error(c, ecode.ErrOrderNotFound)
		return
	} else if err != nil {
		log.Warnf("[v1.order] detail err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, info)
}