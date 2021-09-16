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

// Notify 支付成功回调
// @Summary 支付成功回调
// @Description 支付成功回调
// @Tags 订单
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param req body NotifyParams true "order"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /order/notify [post]
func Notify(c *gin.Context) {
	var req NotifyParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}

	err := service.Svc.OrderPayNotify(c.Request.Context(), app.GetUserID(c), req.Amount, req.PType, req.OrderNo, req.TradeNo, req.TransHash)
	if errors.Is(err, service.ErrOrderNotFound) {
		app.Error(c, ecode.ErrOrderNotFound)
		return
	} else if err != nil {
		log.Warnf("[v1.order] pay notify err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}