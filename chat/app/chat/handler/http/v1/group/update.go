package group

import (
	"chat/app/chat"
	"errors"

	"github.com/gin-gonic/gin"

	"chat/app/chat/ecode"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// Update 更新群组信息
// @Summary 更新群组信息
// @Description 更新群组信息
// @Tags 群组
// @Accept  json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param user body UpdateParams true "The group info"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /group/edit [post]
func Update(c *gin.Context) {
	var req UpdateParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	var err error
	if req.Remark != "" {
		err = chat.Svc.GroupEditRemark(c.Request.Context(), app.GetUInt32UserID(c), req.ID, req.Remark)
	} else if req.Name != "" {
		err = chat.Svc.GroupEditName(c.Request.Context(), app.GetUInt32UserID(c), req.ID, req.Name)
	} else {
		app.Error(c, errno.ErrParamsEmpty)
		return
	}
	if errors.Is(err, chat.ErrGroupNotFound) {
		app.Error(c, ecode.ErrGroupNotFound)
		return
	} else if errors.Is(err, chat.ErrGroupUserNotJoin) {
		app.Error(c, ecode.ErrGroupNotJoin)
		return
	} else if errors.Is(err, chat.ErrGroupDataUnmodified) {
		app.Error(c, ecode.ErrGroupDataUnmodified)
		return
	} else if err != nil {
		log.Warnf("[http.group] update err, %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}

// UpdateNickname 更新群昵称
// @Summary 更新群昵称
// @Description 更新群昵称
// @Tags 群组
// @Accept  json
// @Produce  json
// @Param user body NicknameParams true "The group info"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /group/nickname [post]
func UpdateNickname(c *gin.Context) {
	// Binding the user data.
	var req NicknameParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	err := chat.Svc.GroupEditUserNickname(c.Request.Context(), app.GetUInt32UserID(c), req.ID, req.Nickname)
	if errors.Is(err, chat.ErrGroupNotFound) {
		app.Error(c, ecode.ErrGroupNotFound)
		return
	} else if errors.Is(err, chat.ErrGroupUserNotJoin) {
		app.Error(c, ecode.ErrGroupNotJoin)
		return
	} else if err != nil {
		log.Warnf("[http.group] nickname err, %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}
