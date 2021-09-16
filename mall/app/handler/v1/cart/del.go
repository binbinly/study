package cart

import (
	"github.com/gin-gonic/gin"

	"mall/app/service"
	"mall/pkg/app"
	"mall/pkg/errno"
	"mall/pkg/log"
)

// Del 删除购物车
// @Summary 删除购物车
// @Description 删除购物车
// @Tags 购物车
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param req body DelParams true "del"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /cart/del [post]
func Del(c *gin.Context) {
	var req DelParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}

	err := service.Svc.DelCart(c.Request.Context(), app.GetUserID(c), req.ID)
	if err != nil {
		log.Warnf("[v1.cart] del err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.SuccessNil(c)
}

// Empty 清空购物车
// @Summary 清空购物车
// @Description 清空购物车
// @Tags 购物车
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /cart/empty [get]
func Empty(c *gin.Context) {
	err := service.Svc.EmptyCart(c.Request.Context(), app.GetUserID(c))
	if err != nil {
		log.Warnf("[v1.cart] empty err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}