package order

import (
	"errors"

	"github.com/gin-gonic/gin"

	"mall/app/ecode"
	"mall/app/service"
	"mall/pkg/app"
	"mall/pkg/errno"
	"mall/pkg/log"
)

// Refund 退款
// @Summary 退款
// @Description 退款
// @Tags 订单
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param req body RefundParams true "order"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /order/refund [post]
func Refund(c *gin.Context) {
	var req RefundParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}

	err := service.Svc.OrderRefund(c.Request.Context(), app.GetUserID(c), req.OrderNo, req.Content)
	if errors.Is(err, service.ErrOrderNotFound) {
		app.Error(c, ecode.ErrOrderNotFound)
		return
	} else if err != nil {
		log.Warnf("[v1.order] refund err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}