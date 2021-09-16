package version

import (
	"runtime"

	"gorm.io/gorm"

	"mall/app/model"
	"mall/cmd/migrate/migration"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1629293920500Up)
}

func _1629293920500Up(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		err := tx.Debug().Migrator().AutoMigrate(
			new(model.GoodsModel),
			new(model.GoodsAttrModel),
			new(model.GoodsCategoryModel),
			new(model.GoodsCommentModel),
			new(model.GoodsImageModel),
			new(model.GoodsSkuModel),
			new(model.GoodsSkuAttrModel),
			new(model.SkuAttrModel),
			new(model.SkuAttrValModel),
		)
		if err != nil {
			return err
		}

		return tx.Create(&model.Migration{
			Version: version,
		}).Error
	})
}
