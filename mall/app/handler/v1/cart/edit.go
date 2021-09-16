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

// Edit 更新购物车
// @Summary 更新购物车
// @Description 更新购物车
// @Tags 购物车
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param req body EditParams true "sku"
// @success 0 {object} app.Response{data=model.CartModel} "调用成功结构"
// @Router /cart/edit [post]
func Edit(c *gin.Context) {
	var req EditParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}

	cart, err := service.Svc.EditCart(c.Request.Context(), app.GetUserID(c), req.SkuID, req.Num, req.ID)
	if errors.Is(err, service.ErrGoodsSkuNotFound) {
		app.Error(c, ecode.ErrGoodsSkuNotFound)
		return
	} else if errors.Is(err, service.ErrGoodsSkuNotEdit) {
		app.Error(c, ecode.ErrGoodsSkuNotEdit)
		return
	} else if err != nil {
		log.Warnf("[v1.cart] edit sku err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, cart)
}

// EditNum 更新购物车商品数量
// @Summary 更新购物车商品数量
// @Description 更新购物车商品数量
// @Tags 购物车
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param req body EditNumParams true "num"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /cart/edit_num [post]
func EditNum(c *gin.Context) {
	var req EditNumParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}

	err := service.Svc.EditCartNum(c.Request.Context(), app.GetUserID(c), req.ID, req.Num)
	if errors.Is(err, service.ErrGoodsSkuNotFound) {
		app.Error(c, ecode.ErrGoodsSkuNotFound)
		return
	} else if err != nil {
		log.Warnf("[v1.cart] edit num err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.SuccessNil(c)
}