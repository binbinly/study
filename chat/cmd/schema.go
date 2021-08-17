package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	confcenter "chat/app/center/conf"
	modelcenter "chat/app/center/model"
	"chat/app/chat/conf"
	"chat/app/chat/model"
	"chat/internal/orm"
)

var reset bool
var name string

func init() {
	migrateCmd.Flags().StringVarP(&cfg, "config", "c", "", "config file (default is $ROOT/config/chat/chat.yaml)")
	migrateCmd.Flags().StringVarP(&name, "name", "n", "", "select db name (default is chat)")
	migrateCmd.Flags().BoolVar(&reset, "reset", false, "db migrate reset all")
}

// 运行数据库迁移
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "chat migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		if name == "center" {
			if cfg == "" {
				cfg = "./config/center/center.yaml"
			}
			confcenter.Init(cfg)
			orm.Init(&confcenter.Conf.MySQL)
		} else {
			if cfg == "" {
				cfg = "./config/chat/chat.yaml"
			}
			conf.Init(cfg)
			orm.Init(&conf.Conf.MySQL)
		}
		schema()
	},
}

func schema() {
	if name == "center" {
		if reset {
			centerDown()
		} else {
			centerUp()
		}
		fmt.Println("center schema success")
		return
	}
	if reset {
		chatDown()
	} else {
		chatUp()
	}
	fmt.Println("chat schema success")
}

// 运行迁移
func chatUp() {
	var err error
	if !orm.DB.Migrator().HasTable(&model.ApplyModel{}) {
		err = orm.DB.Migrator().CreateTable(&model.ApplyModel{})
		if err != nil {
			log.Panicf("create table apply err:%v", err)
		}
	}
	if !orm.DB.Migrator().HasTable(&model.FriendModel{}) {
		err = orm.DB.Migrator().CreateTable(&model.FriendModel{})
		if err != nil {
			log.Panicf("create table friend err:%v", err)
		}
	}
	if !orm.DB.Migrator().HasTable(&model.GroupModel{}) {
		err = orm.DB.Migrator().CreateTable(&model.GroupModel{})
		if err != nil {
			log.Panicf("create table group err:%v", err)
		}
	}
	if !orm.DB.Migrator().HasTable(&model.GroupUserModel{}) {
		err = orm.DB.Migrator().CreateTable(&model.GroupUserModel{})
		if err != nil {
			log.Panicf("create table group_user err:%v", err)
		}
	}
	if !orm.DB.Migrator().HasTable(&model.CollectModel{}) {
		err = orm.DB.Migrator().CreateTable(&model.CollectModel{})
		if err != nil {
			log.Panicf("create table collection err:%v", err)
		}
	}
	if !orm.DB.Migrator().HasTable(&model.UserTagModel{}) {
		err = orm.DB.Migrator().CreateTable(&model.UserTagModel{})
		if err != nil {
			log.Panicf("create table user_tag err:%v", err)
		}
	}
	if !orm.DB.Migrator().HasTable(&model.ReportModel{}) {
		err = orm.DB.Migrator().CreateTable(&model.ReportModel{})
		if err != nil {
			log.Panicf("create table report err:%v", err)
		}
	}
	if !orm.DB.Migrator().HasTable(&model.MessageModel{}) {
		err = orm.DB.Migrator().CreateTable(&model.MessageModel{})
		if err != nil {
			log.Panicf("create table message err:%v", err)
		}
	}
	if !orm.DB.Migrator().HasTable(&model.MomentModel{}) {
		err = orm.DB.Migrator().CreateTable(&model.MomentModel{})
		if err != nil {
			log.Panicf("create table moment err:%v", err)
		}
	}
	if !orm.DB.Migrator().HasTable(&model.MomentCommentModel{}) {
		err = orm.DB.Migrator().CreateTable(&model.MomentCommentModel{})
		if err != nil {
			log.Panicf("create table moment comment err:%v", err)
		}
	}
	if !orm.DB.Migrator().HasTable(&model.MomentLikeModel{}) {
		err = orm.DB.Migrator().CreateTable(&model.MomentLikeModel{})
		if err != nil {
			log.Panicf("create table moment like err:%v", err)
		}
	}
	if !orm.DB.Migrator().HasTable(&model.MomentTimelineModel{}) {
		err = orm.DB.Migrator().CreateTable(&model.MomentTimelineModel{})
		if err != nil {
			log.Panicf("create table moment timeline err:%v", err)
		}
	}
	if !orm.DB.Migrator().HasTable(&model.EmoticonModel{}) {
		err = orm.DB.Migrator().CreateTable(&model.EmoticonModel{})
		if err != nil {
			log.Panicf("create table emoticon err:%v", err)
		}
	}
}

// 回滚数据库迁移
func chatDown() {
	var err error
	if orm.DB.Migrator().HasTable(&model.ApplyModel{}) {
		err = orm.DB.Migrator().DropTable(&model.ApplyModel{})
		if err != nil {
			log.Panicf("drop table applu err:%v", err)
		}
	}
	if orm.DB.Migrator().HasTable(&model.GroupModel{}) {
		err = orm.DB.Migrator().DropTable(&model.GroupModel{})
		if err != nil {
			log.Panicf("drop table group err:%v", err)
		}
	}
	if orm.DB.Migrator().HasTable(&model.GroupUserModel{}) {
		err = orm.DB.Migrator().DropTable(&model.GroupUserModel{})
		if err != nil {
			log.Panicf("drop table group_user err:%v", err)
		}
	}
	if orm.DB.Migrator().HasTable(&model.CollectModel{}) {
		err = orm.DB.Migrator().DropTable(&model.CollectModel{})
		if err != nil {
			log.Panicf("drop table collection err:%v", err)
		}
	}
	if orm.DB.Migrator().HasTable(&model.FriendModel{}) {
		err = orm.DB.Migrator().DropTable(&model.FriendModel{})
		if err != nil {
			log.Panicf("drop table friend err:%v", err)
		}
	}
	if orm.DB.Migrator().HasTable(&model.UserTagModel{}) {
		err = orm.DB.Migrator().DropTable(&model.UserTagModel{})
		if err != nil {
			log.Panicf("drop table user_tag err:%v", err)
		}
	}
	if orm.DB.Migrator().HasTable(&model.ReportModel{}) {
		err = orm.DB.Migrator().DropTable(&model.ReportModel{})
		if err != nil {
			log.Panicf("drop table report err:%v", err)
		}
	}
	if orm.DB.Migrator().HasTable(&model.MessageModel{}) {
		err = orm.DB.Migrator().DropTable(&model.MessageModel{})
		if err != nil {
			log.Panicf("drop table message err:%v", err)
		}
	}
	if orm.DB.Migrator().HasTable(&model.MomentModel{}) {
		err = orm.DB.Migrator().DropTable(&model.MomentModel{})
		if err != nil {
			log.Panicf("drop table moment err:%v", err)
		}
	}
	if orm.DB.Migrator().HasTable(&model.MomentCommentModel{}) {
		err = orm.DB.Migrator().DropTable(&model.MomentCommentModel{})
		if err != nil {
			log.Panicf("drop table moment comment err:%v", err)
		}
	}
	if orm.DB.Migrator().HasTable(&model.MomentLikeModel{}) {
		err = orm.DB.Migrator().DropTable(&model.MomentLikeModel{})
		if err != nil {
			log.Panicf("drop table moment like err:%v", err)
		}
	}
	if orm.DB.Migrator().HasTable(&model.MomentTimelineModel{}) {
		err = orm.DB.Migrator().DropTable(&model.MomentTimelineModel{})
		if err != nil {
			log.Panicf("drop table moment timeline err:%v", err)
		}
	}
	if orm.DB.Migrator().HasTable(&model.EmoticonModel{}) {
		err = orm.DB.Migrator().DropTable(&model.EmoticonModel{})
		if err != nil {
			log.Panicf("drop table emoticon err:%v", err)
		}
	}
}

func centerUp() {
	var err error
	if !orm.DB.Migrator().HasTable(&modelcenter.UserModel{}) {
		err = orm.DB.Migrator().CreateTable(&modelcenter.UserModel{})
		if err != nil {
			log.Panicf("create center table user err:%v", err)
		}
	}
}

func centerDown() {
	var err error
	if orm.DB.Migrator().HasTable(&modelcenter.UserModel{}) {
		err = orm.DB.Migrator().DropTable(&modelcenter.UserModel{})
		if err != nil {
			log.Panicf("drop table emoticon err:%v", err)
		}
	}
}
