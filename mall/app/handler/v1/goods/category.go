package goods

import (
	"github.com/gin-gonic/gin"

	"mall/app/service"
	"mall/pkg/app"
	"mall/pkg/errno"
	"mall/pkg/log"
)

// CategoryAll 所有商品分类
// @Summary 所有商品分类
// @Description 所有商品分类 - 树形结构
// @Tags 商品
// @Accept json
// @Produce json
// @success 0 {object} app.Response{data=[]model.GoodsCategory} "调用成功结构"
// @Router /goods/category [get]
func CategoryAll(c *gin.Context) {
	list, err := service.Svc.CategoryTree(c.Request.Context())
	if err != nil {
		log.Warnf("[v1.category] list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}
