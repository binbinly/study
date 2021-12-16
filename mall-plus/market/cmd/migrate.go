package cmd

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"common/orm"
	"market/conf"
	"market/model"
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
			new(model.ConfigModel),
			new(model.AppNoticeModel),
			new(model.AppSettingModel),
			new(model.CouponModel),
			new(model.CouponMemberModel),
			new(model.CouponRelCatModel),
			new(model.CouponRelSpuModel),
			new(model.HomeAdvModel),
			new(model.HomeSubjectModel),
			new(model.HomeSubjectSpuModel),
			new(model.MemberPriceModel),
			new(model.SeckillActivityModel),
			new(model.SeckillSkuModel),
			new(model.SeckillSessionModel),
			new(model.SeckillSkuNoticeModel),
			new(model.SkuFullReductionModel),
			new(model.SkuLadderModel),
			new(model.SpuBoundsModel),
		)
		if err != nil {
			return err
		}
		return nil
	})
}
