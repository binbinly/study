package friend

import (
	"chat/app/chat"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// List 好友列表
// @Summary 好友列表
// @Description 好友列表
// @Tags 好友
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @success 0 {object} app.Response{data=[]model.UserBase} "调用成功结构"
// @Router /friend/list [get]
func List(c *gin.Context) {
	list, err := chat.Svc.FriendMyAll(c.Request.Context(), app.GetUInt32UserID(c))
	if err != nil {
		log.Warnf("[http.friend] list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}

// TagList 标签好友列表
// @Summary 标签好友列表
// @Description 标签好友列表
// @Tags 好友
// @Accept json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Param id query int true "标签ID"
// @success 0 {object} app.Response{data=[]model.UserBase} "调用成功结构"
// @Router /friend/tag_list [get]
func TagList(c *gin.Context) {
	tagID := cast.ToUint32(c.Query("id"))
	if tagID == 0 {
		app.Error(c, errno.ErrBind)
		return
	}

	list, err := chat.Svc.FriendMyListByTagID(c.Request.Context(), app.GetUInt32UserID(c), tagID)
	if err != nil {
		log.Warnf("[http.friend] tag list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}
