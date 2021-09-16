package cart

import (
	"github.com/gin-gonic/gin"

	"mall/app/service"
	"mall/pkg/app"
	"mall/pkg/errno"
	"mall/pkg/log"
)

// List 我的购物车
// @Summary 我的购物车
// @Description 我的购物车
// @Tags 购物车
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @success 0 {object} app.Response{data=[]model.CartModel} "调用成功结构"
// @Router /cart/list [get]
func List(c *gin.Context) {
	list, err := service.Svc.CartList(c.Request.Context(), app.GetUserID(c))
	if err != nil {
		log.Warnf("[v1.cart] list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}
