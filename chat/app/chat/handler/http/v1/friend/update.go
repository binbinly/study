package friend

import (
	"chat/app/chat"
	"errors"

	"github.com/gin-gonic/gin"

	"chat/app/chat/ecode"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// Black 加入/移除黑名单
// @Summary 加入/移除黑名单
// @Description 加入/移除黑名单
// @Tags 好友
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body BlackParams true "black"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /friend/black [post]
func Black(c *gin.Context) {
	var req BlackParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	userID := app.GetUInt32UserID(c)
	if userID == req.UserID {
		app.Error(c, ecode.ErrUserNoSelf)
		return
	}
	err := chat.Svc.FriendSetBlack(c.Request.Context(), userID, req.UserID, req.Black)
	if errors.Is(err, chat.ErrFriendNotRecord) {
		app.Error(c, ecode.ErrFriendNotFound)
		return
	} else if err != nil {
		log.Warnf("[http.friend] black err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}

// Star 加入/移除星标
// @Summary 加入/移除星标
// @Description 加入/移除星标
// @Tags 好友
// @Accept json
// @Produce  json
// @Param Authorization header string true "Authentication header"
// @Param req body StarParams true "star"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /friend/star [post]
func Star(c *gin.Context) {
	var req StarParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	userID := app.GetUInt32UserID(c)
	if userID == req.UserID {
		app.Error(c, ecode.ErrUserNoSelf)
		return
	}
	err := chat.Svc.FriendSetStar(c.Request.Context(), userID, req.UserID, req.Star)
	if errors.Is(err, chat.ErrFriendNotRecord) {
		app.Error(c, ecode.ErrFriendNotFound)
		return
	} else if err != nil {
		log.Warnf("[http.friend] star err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}

// Remark 设置备注标签
// @Summary 设置备注标签
// @Description 设置备注标签
// @Tags 好友
// @Accept json
// @Produce  json
// @Param Authorization header string true "Authentication header"
// @Param req body RemarkParams true "remark"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /friend/remark [post]
func Remark(c *gin.Context) {
	var req RemarkParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	userID := app.GetUInt32UserID(c)
	if userID == req.UserID {
		app.Error(c, ecode.ErrUserNoSelf)
		return
	}
	err := chat.Svc.FriendSetRemarkTag(c.Request.Context(), userID, req.UserID, req.Nickname, req.Tags)
	if errors.Is(err, chat.ErrFriendNotRecord) {
		app.Error(c, ecode.ErrFriendNotFound)
		return
	} else if err != nil {
		log.Warnf("[http.friend] remark err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}

// Auth 设置朋友圈权限
// @Summary 设置朋友圈权限
// @Description 设置朋友圈权限
// @Tags 好友
// @Accept json
// @Produce  json
// @Param Authorization header string true "Authentication header"
// @Param req body AuthParams true "auth"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /friend/auth [post]
func Auth(c *gin.Context) {
	var req AuthParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	userID := app.GetUInt32UserID(c)
	if userID == req.UserID {
		app.Error(c, ecode.ErrUserNoSelf)
		return
	}
	err := chat.Svc.FriendSetMomentAuth(c.Request.Context(), userID, req.UserID, req.LookMe, req.LookHim)
	if errors.Is(err, chat.ErrFriendNotRecord) {
		app.Error(c, ecode.ErrFriendNotFound)
		return
	} else if err != nil {
		log.Warnf("[http.friend] auth err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}
