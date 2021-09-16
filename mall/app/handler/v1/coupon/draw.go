package coupon

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

// Draw 领取优惠券
// @Summary 领取优惠券
// @Description 领取优惠券
// @Tags 优惠券
// @Accept json
// @Produce json
// @Param id body int true "优惠券id"
// @Param Token header string true "用户令牌"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /coupon/draw [get]
func Draw(c *gin.Context) {
	id := cast.ToInt(c.Query("id"))
	if id <= 0 {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	err := service.Svc.CouponDraw(c.Request.Context(), app.GetUserID(c), id)
	if errors.Is(err, service.ErrCouponNotFound) {
		app.Error(c, ecode.ErrCouponNotFound)
		return
	} else if errors.Is(err, service.ErrCouponNoNum) {
		app.Error(c, ecode.ErrCouponNoNum)
		return
	} else if errors.Is(err, service.ErrCouponReceived) {
		app.Error(c, ecode.ErrCouponReceived)
		return
	} else if err != nil {
		log.Warnf("[v1.coupon] draw err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}