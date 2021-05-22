package friend

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"chat/app/logic/service"
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
// @Success 200 {string} json "{"code":0,"message":"OK","data":null}"
// @Router /friend/list [get]
func List(c *gin.Context) {
	list, err := service.Svc.FriendMyAll(c, app.GetUInt32UserId(c))
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
// @Success 200 {string} json "{"code":0,"message":"OK","data":null}"
// @Router /friend/tag_list [get]
func TagList(c *gin.Context) {
	tagId := cast.ToUint32(c.Query("id"))
	if tagId == 0 {
		app.Error(c, errno.ErrBind)
		return
	}

	list, err := service.Svc.FriendMyListByTagId(c, app.GetUInt32UserId(c), tagId)
	if err != nil {
		log.Warnf("[http.friend] tag list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}
