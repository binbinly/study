package cart

import (
	"errors"

	"github.com/gin-gonic/gin"

	"mall/app/ecode"
	"mall/app/service"
	"mall/pkg/app"
	"mall/pkg/errno"
	"mall/pkg/log"
)

// Add 添加购物车
// @Summary 添加购物车
// @Description 添加购物车
// @Tags 购物车
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param req body AddParams true "sku"
// @success 0 {object} app.Response{data=model.CartModel} "调用成功结构"
// @Router /cart/add [post]
func Add(c *gin.Context) {
	var req AddParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}

	cart, err := service.Svc.AddCart(c.Request.Context(), app.GetUserID(c), req.GoodsID, req.SkuID, req.Num)
	if errors.Is(err, service.ErrGoodsSkuNotFound) {
		app.Error(c, ecode.ErrGoodsSkuNotFound)
		return
	} else if errors.Is(err, service.ErrGoodsNotFound) {
		app.Error(c, ecode.ErrGoodsNotFound)
		return
	} else if err != nil {
		log.Warnf("[v1.cart] edit sku err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, cart)
}