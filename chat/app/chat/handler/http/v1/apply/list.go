package apply

import (
	"chat/app/chat"
	"chat/app/chat/handler/http"
	"github.com/gin-gonic/gin"

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
// @success 0 {object} app.Response{data=[]model.ApplyList} "调用成功结构"
// @Router /apply/list [get]
func List(c *gin.Context) {
	userID := app.GetUInt32UserID(c)

	list, err := chat.Svc.ApplyMyList(c.Request.Context(), userID, http.GetPageOffset(c))
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
	userID := app.GetUInt32UserID(c)
	count, err := chat.Svc.ApplyPendingCount(c.Request.Context(), userID)
	if err != nil {
		log.Warnf("[http.count] err: %v", err)
		app.Success(c, 0)
		return
	}
	app.Success(c, count)
}
