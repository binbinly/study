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

// Cancel 取消订单
// @Summary 取消订单
// @Description 取消订单
// @Tags 订单
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param req body NoParams true "del"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /order/cancel [post]
func Cancel(c *gin.Context) {
	var req NoParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}

	err := service.Svc.OrderCancel(c.Request.Context(), app.GetUserID(c), req.OrderNo)
	if errors.Is(err, service.ErrOrderNotFound) {
		app.Error(c, ecode.ErrOrderNotFound)
		return
	} else if err != nil {
		log.Warnf("[v1.order] cancel err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}
