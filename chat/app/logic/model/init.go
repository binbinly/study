package model

import (
	"gorm.io/gorm"

	"chat/pkg/database/orm"
)

// DB 数据库全局变量
var DB *gorm.DB

// Init 初始化数据库
func Init(cfg *orm.Config) *gorm.DB {
	DB = orm.NewMySQL(cfg)
	return DB
}

// GetDB 返回默认的数据库
func GetDB() *gorm.DB {
	return DB
}

//CloseDB 关闭DB连接
func CloseDB()  {
	orm.CloseDB()
}