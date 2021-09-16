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

// Receipt 确认收货
// @Summary 确认收货
// @Description 确认收货
// @Tags 订单
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param req body NoParams true "order"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /order/receipt [post]
func Receipt(c *gin.Context) {
	var req NoParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}

	err := service.Svc.OrderConfirmReceipt(c.Request.Context(), app.GetUserID(c), req.OrderNo)
	if errors.Is(err, service.ErrOrderNotFound) {
		app.Error(c, ecode.ErrOrderNotFound)
		return
	} else if err != nil {
		log.Warnf("[v1.order] receipt err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}