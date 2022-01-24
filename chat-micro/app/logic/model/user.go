package model

import (
	"time"

	"gorm.io/gorm"

	"chat-micro/internal/orm"
	"chat-micro/pkg/crypt/auth"
)

const (
	//UserStatusNormal 状态 - 正常
	UserStatusNormal = iota + 1
	//UserStatusDisable 状态 - 禁用
	UserStatusDisable
)

//UserModel 会员模型
type UserModel struct {
	orm.PriID
	Username   string    `json:"username" gorm:"column:username;not null;uniqueIndex;type:varchar(64);comment:用户名"`
	Nickname   string    `json:"nickname" gorm:"column:nickname;not null;type:varchar(64);default:'';comment:昵称"`
	Password   string    `json:"password" gorm:"column:password;not null;type:varchar(255);comment:密码"`
	Phone      int64     `gorm:"column:phone;not null;uniqueIndex;comment:手机号" json:"phone"`
	Email      string    `gorm:"column:email;not null;type:varchar(60);default:'';comment:邮箱" json:"email"`
	Avatar     string    `gorm:"column:avatar;not null;type:varchar(128);default:'';comment:头像" json:"avatar"`
	Gender     int8      `gorm:"column:gender;not null;default:1;comment:性别" json:"gender"`
	Birth      time.Time `json:"birth" gorm:"column:birth;type:date;comment:生日"`
	Area       string    `json:"area" gorm:"column:area;not null;type:varchar(255);default:'';comment:城市"`
	Job        string    `json:"job" gorm:"column:job;not null;type:varchar(255);default:'';comment:职业"`
	SourceType int8      `json:"source_type" gorm:"column:source_type;not null;default:0;comment:用户来源"`
	Status     int8      `gorm:"column:status;not null;default:1;comment:状态" json:"status"`
	Sign       string    `gorm:"column:sign;not null;type:varchar(255);default:'';comment:签名" json:"sign"`
	orm.UpdateTime
}

//User 对外用户结构
type User struct {
	ID       uint32 `json:"id"`
	Phone    int64  `json:"phone"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Sign     string `json:"sign"`
	Area     string `json:"area"`
	Gender   int8   `json:"gender"`
}

//UserBase 对外基础用户结构
type UserBase struct {
	ID     uint32 `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
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
