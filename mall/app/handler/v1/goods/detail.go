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

// Detail 商品详情
// @Summary 商品详情
// @Description 商品详情
// @Tags 商品
// @Accept json
// @Produce json
// @Param id body int true "商品id"
// @success 0 {object} app.Response{data=[]model.GoodsDetail} "调用成功结构"
// @Router /goods/detail [get]
func Detail(c *gin.Context) {
	//商品分类
	id := cast.ToInt(c.Query("id"))
	if id == 0 {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	info, err := service.Svc.GoodsDetail(c.Request.Context(), id)
	if errors.Is(err, service.ErrGoodsNotFound) {
		app.Error(c, ecode.ErrGoodsNotFound)
		return
	} else if err != nil {
		log.Warnf("[v1.goods] detail err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, info)
}
