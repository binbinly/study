package user

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

// UpdateParams 修改用户信息
type UpdateParams struct {
	Avatar   string `json:"avatar" binding:"omitempty,url" example:"http://example"` // 头像
	Nickname string `json:"nickname" binding:"omitempty,max=30" example:"test"`      // 昵称
	Sign     string `json:"sign" binding:"omitempty,max=90" example:"test"`          // 签名
}

// AddressAddParams 添加收货地址
type AddressAddParams struct {
	Name      string `json:"name" binding:"required,max=30"`
	Phone     string `json:"phone" binding:"required,max=11"`
	Province  string `json:"province" binding:"required"`
	City      string `json:"city" binding:"required"`
	County    string `json:"county" binding:"required"`
	AreaCode  int    `json:"area_code" binding:"required,numeric"`
	Detail    string `json:"detail" binding:"required,max=120"`
	IsDefault int8   `json:"is_default" binding:"omitempty,numeric"`
}

//AddressEditParams 修改收货地址
type AddressEditParams struct {
	ID        int    `json:"id" binding:"required,numeric"`
	Name      string `json:"name" binding:"required,max=30"`
	Phone     string `json:"phone" binding:"required,max=11"`
	Province  string `json:"province" binding:"required"`
	City      string `json:"city" binding:"required"`
	County    string `json:"county" binding:"required"`
	AreaCode  int    `json:"area_code" binding:"required,numeric"`
	Detail    string `json:"detail" binding:"required,max=120"`
	IsDefault int8   `json:"is_default" binding:"omitempty,numeric"`
}
