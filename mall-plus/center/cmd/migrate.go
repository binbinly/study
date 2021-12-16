package cmd

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"center/conf"
	"center/model"
	"common/orm"
)

//Migrate 数据迁移
func Migrate() {
	//3. 初始化数据库链接
	orm.Init(&conf.Conf.MySQL)
	//4. 数据库迁移
	fmt.Println("数据库迁移开始")
	err := migrateUp(orm.DB)
	if err != nil {
		log.Fatalf("db err:%v", err)
	}
	fmt.Println(`数据库基础数据初始化成功`)
}

//migrateUp 数据迁移
func migrateUp(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		err := tx.Debug().Migrator().AutoMigrate(
			new(model.UserModel),
		)
		if err != nil {
			return err
		}
		return nil
	})
}
