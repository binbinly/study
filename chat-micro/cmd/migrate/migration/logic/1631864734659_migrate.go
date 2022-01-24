package logic

import (
	"runtime"

	"gorm.io/gorm"

	"chat-micro/app/logic/model"
	"chat-micro/cmd/migrate/migration"
	"chat-micro/internal/orm"
)

func init1631864734659() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1631864734659Up)
}

func _1631864734659Up(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		err := tx.Debug().Migrator().AutoMigrate(
			new(model.ApplyModel),
			new(model.FriendModel),
			new(model.GroupModel),
			new(model.GroupUserModel),
			new(model.CollectModel),
			new(model.UserTagModel),
			new(model.ReportModel),
			new(model.MessageModel),
			new(model.MomentModel),
			new(model.MomentCommentModel),
			new(model.MomentLikeModel),
			new(model.MomentTimelineModel),
			new(model.EmoticonModel),
			new(model.UserModel),
		)
		if err != nil {
			return err
		}

		return tx.Create(&orm.Migration{
			Version: version,
		}).Error
	})
}
