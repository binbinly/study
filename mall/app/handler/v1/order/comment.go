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

// Comment 评价
// @Summary 评价
// @Description 评价
// @Tags 订单
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param req body CommentParams true "order"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /order/comment [post]
func Comment(c *gin.Context) {
	var req CommentParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}

	err := service.Svc.OrderComment(c.Request.Context(), app.GetUserID(c), req.Rate, req.OrderNo, req.Content, req.Ids)
	if errors.Is(err, service.ErrOrderNotFound) {
		app.Error(c, ecode.ErrOrderNotFound)
		return
	} else if err != nil {
		log.Warnf("[v1.order] comment err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}