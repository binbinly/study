package apply

import (
	"github.com/gin-gonic/gin"

	"chat/app/logic/service"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// List 申请列表
// @Summary 我的申请列表
// @Description 我的申请列表
// @Tags 好友申请
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param p query int false "页码"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6Ik"}}"
// @Router /apply/list [get]
func List(c *gin.Context) {
	userId := app.GetUInt32UserId(c)

	list, err := service.Svc.ApplyMyList(c, userId, app.GetPageOffset(c))
	if err != nil {
		log.Warnf("[http.apply] list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}

// Count 申请数量
// @Summary 待处理申请数量
// @Description 待处理申请数量
// @Tags 好友申请
// @Produce  json
// @Param Token header string true "用户令牌"
// @Success 200 {string} json "{"code":0,"message":"OK","data":1}"
// @Router /apply/count [get]
func Count(c *gin.Context) {
	userId := app.GetUInt32UserId(c)
	count, err := service.Svc.ApplyPendingCount(c, userId)
	if err != nil {
		log.Warnf("[http.count] err: %v", err)
		app.Success(c, 0)
		return
	}
	app.Success(c, count)
}
