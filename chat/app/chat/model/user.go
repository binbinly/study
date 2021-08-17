package model

// UserBase 用户基础信息
type UserBase struct {
	ID     uint32 `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// UserEs 存入es中结构
type UserEs struct {
	ID       uint32 `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
}
