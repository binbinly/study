package user

// SendCodeParams 发送短信验证码
type SendCodeParams struct {
	Phone string `json:"phone" binding:"required" example:"13333333333"` //手机号
}

// SearchParams 搜索
type SearchParams struct {
	Keyword string `json:"keyword" binding:"required,max=18" example:"test"` //关键字
}

// RegisterParams 注册
type RegisterParams struct {
	Phone           string `json:"phone" form:"phone" binding:"required" example:"13333333333"`                                   //手机号
	Username        string `json:"username" form:"username" binding:"required,min=4,max=18" example:"test"`                       //用户名
	Password        string `json:"password" form:"password" binding:"required,min=6,max=20" example:"123456"`                     //密码
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required,eqfield=Password" example:"123456"` //确认密码
}

// LoginParams 默认登录方式-用户名
type LoginParams struct {
	Username string `json:"username" form:"username" binding:"required,min=4,max=18" example:"test"`   //用户名
	Password string `json:"password" form:"password" binding:"required,min=6,max=20" example:"123456"` //密码
}

// PhoneLoginParams 手机号登录
type PhoneLoginParams struct {
	Phone      int64  `json:"phone" form:"phone" binding:"required" example:"13333333333"`        // 手机号
	VerifyCode string `json:"verify_code" form:"verify_code" binding:"required" example:"888888"` //验证码
}

// UpdateParams 修改用户信息
type UpdateParams struct {
	Avatar   string `json:"avatar" binding:"omitempty,url" example:"http://example"` // 头像
	Nickname string `json:"nickname" binding:"omitempty,max=30" example:"test"`      // 昵称
	Sign     string `json:"sign" binding:"omitempty,max=90" example:"test"`          // 签名
}

// ReportParams 好友举报
type ReportParams struct {
	UserId   uint32 `json:"user_id" binding:"required,numeric" example:"1"` //用户ID
	Type     int8   `json:"type" binding:"oneof=1 2" example:"1"`           // 1=用户，2=群组
	Content  string `json:"content" binding:"required" example:"test"`      // 举报内容
	Category string `json:"category" binding:"required" example:"分类"`       // 举报分类
}
