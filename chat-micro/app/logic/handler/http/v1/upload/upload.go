package upload

import (
	"chat-micro/app/logic/service"
	"chat-micro/pkg/app"
	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
	"github.com/gin-gonic/gin"
)

// SignUrl 获取上传url
// @Summary 获取文件上传url
// @Description 文件上传
// @Tags 文件上传
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body UrlParams true "report"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{"url":"http://example"}}"
// @Router /upload/url [post]
func SignUrl(c *gin.Context) {
	var req UrlParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	url, err := service.Svc.GetUploadSIgnUrl(c.Request.Context(), req.Name)
	if err != nil {
		logger.Warnf("[http.upload] sign url err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, url)
}
