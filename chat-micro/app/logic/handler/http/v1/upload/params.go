package upload

// UrlParams 上传url
type UrlParams struct {
	Name string `json:"name" binding:"required" example:"1.jpg"` // 文件名
}
