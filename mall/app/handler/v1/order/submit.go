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

// Submit 提交订单
// @Summary 提交订单
// @Description 提交订单
// @Tags 订单
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param req body SubmitParams true "order"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /order/submit [post]
func Submit(c *gin.Context) {
	var req SubmitParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}

	id, err := service.Svc.SubmitOrder(c.Request.Context(), app.GetUserID(c), req.Ids, req.AddressID, req.CouponID, req.Remark)
	if errors.Is(err, service.ErrGoodsEmpty) {
		app.Error(c, ecode.ErrGoodsEmpty)
		return
	} else if errors.Is(err, service.ErrCouponNotUse) {
		app.Error(c, ecode.ErrCouponNotUse)
		return
	} else if errors.Is(err, service.ErrUserAddressNotFound) {
		app.Error(c, ecode.ErrUserAddressNotFound)
		return
	} else if errors.Is(err, service.ErrCouponNotFound) {
		app.Error(c, ecode.ErrCouponNotFound)
		return
	} else if err != nil {
		log.Warnf("[v1.order] submit err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, id)
}

// GoodsSubmit 商品提交订单
// @Summary 商品提交订单
// @Description 商品直接提交订单
// @Tags 订单
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param req body SubmitGoodsParams true "order"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /order/goods_submit [post]
func GoodsSubmit(c *gin.Context) {
	var req SubmitGoodsParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}

	id, err := service.Svc.SubmitOrderGoods(c.Request.Context(), app.GetUserID(c), req.GoodsID, req.SkuID, req.Num, req.AddressID, req.CouponID, req.Remark)
	if errors.Is(err, service.ErrGoodsEmpty) {
		app.Error(c, ecode.ErrGoodsEmpty)
		return
	} else if errors.Is(err, service.ErrCouponNotUse) {
		app.Error(c, ecode.ErrCouponNotUse)
		return
	} else if errors.Is(err, service.ErrUserAddressNotFound) {
		app.Error(c, ecode.ErrUserAddressNotFound)
		return
	} else if errors.Is(err, service.ErrCouponNotFound) {
		app.Error(c, ecode.ErrCouponNotFound)
		return
	} else if errors.Is(err, service.ErrGoodsNotFound) {
		app.Error(c, ecode.ErrGoodsNotFound)
		return
	} else if err != nil {
		log.Warnf("[v1.order] goods submit err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, id)
}