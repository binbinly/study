package model

import (
	"chat/pkg/crypt/auth"
	"gorm.io/gorm"
)

const (
	//StatusNormal 状态 - 正常
	StatusNormal = iota
	//StatusDisable 状态 - 禁用
	StatusDisable
)

//UserModel 用户模型
type UserModel struct {
	PriID
	Username string `json:"username" gorm:"column:username;not null;uniqueIndex;type:varchar(60);comment:用户名"`
	Nickname string `json:"nickname" gorm:"column:nickname;not null;type:varchar(60);default:'';comment:昵称"`
	Password string `json:"password" gorm:"column:password;not null;type:varchar(255);comment:密码"`
	Phone    int64  `gorm:"column:phone;not null;uniqueIndex;comment:手机号" json:"phone"`
	Email    string `gorm:"column:email;not null;type:varchar(60);default:'';comment:邮箱" json:"email"`
	Avatar   string `gorm:"column:avatar;not null;type:varchar(128);default:'';comment:头像" json:"avatar"`
	Status   int8   `gorm:"column:status;not null;default:0;comment:状态" json:"status"`
	UpdateTime
}

// UserInfo 对外暴露的用户信息结构体
type UserInfo struct {
	ID       uint32 `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

// UserBase 用户基础信息
type UserBase struct {
	ID     uint32 `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// UserToken 登录后令牌信息
type UserToken struct {
	User  *UserInfo `json:"user"`
	Token string    `json:"token"`
}

// TableName 表名
func (u *UserModel) TableName() string {
	return "user"
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *UserModel) Compare(pwd string) (err error) {
	return auth.Compare(u.Password, pwd)
}

// BeforeSave 保存前回调
func (u *UserModel) BeforeSave(tx *gorm.DB) (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return err
}
