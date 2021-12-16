package orm

import (
	"log"
	"os"

	"gorm.io/gorm"

	"pkg/database/mysql"
)

// DB 数据库全局变量
var DB *gorm.DB

// Init 初始化数据库
func Init(cfg *mysql.Config) *gorm.DB {
	DB = mysql.NewMySQL(cfg)
	return DB
}

// InitTest 初始化测试数据库
func InitTest(dbName string) *gorm.DB {
	addr := os.Getenv("MYSQL_ADDR")
	if addr == "" {
		addr = "192.168.8.76:3306"
	}
	username := os.Getenv("MYSQL_USERNAME")
	if username == "" {
		username = "root"
	}
	password := os.Getenv("MYSQL_PASSWORD")
	if password == "" {
		password = "123456"
	}
	log.Printf("test mysql addr: %v, username: %v, password: %v", addr, username, password)
	DB = mysql.NewMySQL(&mysql.Config{
		Name:            dbName,
		Addr:            addr,
		UserName:        username,
		Password:        password,
		TablePrefix:     "",
		Debug:           true,
		MaxIdleConn:     2,
		MaxOpenConn:     5,
		ConnMaxLifeTime: 60,
	})
	return DB
}

// GetDB 返回默认的数据库
func GetDB() *gorm.DB {
	return DB
}

//CloseDB 关闭DB连接
func CloseDB() error {
	return mysql.CloseDB()
}
