package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentracing"
)

var mysqlDB *sql.DB

// Config mysql config
type Config struct {
	Name            string
	Addr            string
	UserName        string
	Password        string
	TablePrefix     string
	ShowLog         bool
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifeTime time.Duration
}

// NewMySQL 链接数据库，生成数据库实例
func NewMySQL(c *Config) (db *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		c.UserName,
		c.Password,
		c.Addr,
		c.Name,
		true,
		//"Asia/Shanghai"),
		"Local")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 2 * time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Warn,     // Log level
			Colorful:      false,           // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true, //禁用自动创建数据库外键约束
		PrepareStmt:                              true, //PreparedStmt 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.TablePrefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,          // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		log.Panicf("open mysql failed. database name: %s, err: %+v", c.Name, err)
	}
	mysqlDB, err = db.DB()
	if err != nil {
		log.Panicf("Database failed err: %+v", err)
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	mysqlDB.SetMaxIdleConns(c.MaxIdleConn)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	mysqlDB.SetMaxOpenConns(c.MaxOpenConn)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	mysqlDB.SetConnMaxLifetime(c.ConnMaxLifeTime)

	err = db.Use(gormopentracing.New())
	if err != nil {
		log.Panicf("using gorm opentracing, err: %+v", err)
	}

	return db
}

// CloseDB closes database connection (unnecessary)
func CloseDB() error {
	if mysqlDB == nil {
		return nil
	}
	return mysqlDB.Close()
}
