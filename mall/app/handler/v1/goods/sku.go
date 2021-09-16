package goods

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

// Sku 商品sku
// @Summary 商品sku
// @Description 商品sku
// @Tags 商品
// @Accept json
// @Produce json
// @Param id body int true "商品id"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{"id":1,"stock":"0","sku_many":1,"skus":[],"sku_attrs":[]}}"
// @Router /goods/sku [get]
func Sku(c *gin.Context) {
	//商品分类
	id := cast.ToInt(c.Query("id"))
	if id == 0 {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	info, err := service.Svc.GoodsSku(c.Request.Context(), id)
	if errors.Is(err, service.ErrGoodsNotFound) {
		app.Error(c, ecode.ErrGoodsNotFound)
		return
	} else if err != nil {
		log.Warnf("[v1.goods] sku err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, info)
}