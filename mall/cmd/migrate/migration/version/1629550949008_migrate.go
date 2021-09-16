package version

import (
	"runtime"

	"gorm.io/gorm"

	"mall/app/model"
	"mall/cmd/migrate/migration"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1629550949008Up)
}

func _1629550949008Up(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		err := tx.Debug().Migrator().AutoMigrate(
			new(model.AppNoticeModel),
			new(model.AppSettingModel),
			new(model.AppConfigModel),
			new(model.AreaModel),
			new(model.CouponModel),
			new(model.CouponUserModel),
			new(model.OrderModel),
			new(model.OrderAddressModel),
			new(model.OrderGoodsModel),
			new(model.OrderRefundModel),
			new(model.UserModel),
			new(model.UserAddressModel),
			new(model.UserLevelModel),
		)
		if err != nil {
			return err
		}

		return tx.Create(&model.Migration{
			Version: version,
		}).Error
	})
}
