package coupon

import (
	"github.com/gin-gonic/gin"

	"mall/app/constvar"
	"mall/app/handler"
	"mall/app/service"
	"mall/pkg/app"
	"mall/pkg/errno"
	"mall/pkg/log"
)

// List 优惠券列表
// @Summary 优惠券列表
// @Description 优惠券列表
// @Tags 优惠券
// @Accept json
// @Produce json
// @Param p body int false "分页"
// @Param Token header string true "用户令牌"
// @success 0 {object} app.Response{data=[]model.Coupon} "调用成功结构"
// @Router /coupon/list [get]
func List(c *gin.Context) {
	list, err := service.Svc.GetCouponList(c.Request.Context(), app.GetUserID(c), handler.GetPageOffset(c), constvar.DefaultLimit)
	if err != nil {
		log.Warnf("[v1.coupon] list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}

// My 我的优惠券
// @Summary 我的优惠券
// @Description 我的优惠券
// @Tags 优惠券
// @Accept json
// @Produce json
// @Param p body int false "分页"
// @Param Token header string true "用户令牌"
// @success 0 {object} app.Response{data=[]model.Coupon} "调用成功结构"
// @Router /coupon/my [get]
func My(c *gin.Context) {
	list, err := service.Svc.GetMyCouponList(c.Request.Context(), app.GetUserID(c), handler.GetPageOffset(c), constvar.DefaultLimit)
	if err != nil {
		log.Warnf("[v1.coupon] my list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}