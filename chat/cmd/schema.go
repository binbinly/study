package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"chat/app/logic/conf"
	"chat/app/logic/model"
)

var reset bool
var emo bool

func init() {
	migrateCmd.Flags().StringVarP(&cfg, "config", "c", "", "config file (default is $ROOT/config/logic.yaml)")
	migrateCmd.Flags().BoolVar(&reset, "reset", false, "db migrate reset all")
	migrateCmd.Flags().BoolVar(&emo, "emo", false, "sync emoticon to db")
}

// 运行数据库迁移
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "chat migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		if cfg == "" {
			cfg = "./config/logic.yaml"
		}
		conf.Init(cfg)
		schema()
	},
}

func schema() {
	model.Init(&conf.Conf.MySQL)
	if reset {
		down()
	} else if emo {
		fmt.Println("sync emoticon start")
		err := SyncBQB()
		if err != nil {
			fmt.Printf("err:%v\n", err)
		}
	} else {
		up()
	}
	fmt.Println("schema success")
}

// 运行迁移
func up() {
	var err error
	if !model.DB.Migrator().HasTable(&model.UserModel{}) {
		err = model.DB.Migrator().CreateTable(&model.UserModel{})
		if err != nil {
			log.Panicf("create table user err:%v", err)
		}
	}
	if !model.DB.Migrator().HasTable(&model.ApplyModel{}) {
		err = model.DB.Migrator().CreateTable(&model.ApplyModel{})
		if err != nil {
			log.Panicf("create table apply err:%v", err)
		}
	}
	if !model.DB.Migrator().HasTable(&model.FriendModel{}) {
		err = model.DB.Migrator().CreateTable(&model.FriendModel{})
		if err != nil {
			log.Panicf("create table friend err:%v", err)
		}
	}
	if !model.DB.Migrator().HasTable(&model.GroupModel{}) {
		err = model.DB.Migrator().CreateTable(&model.GroupModel{})
		if err != nil {
			log.Panicf("create table group err:%v", err)
		}
	}
	if !model.DB.Migrator().HasTable(&model.GroupUserModel{}) {
		err = model.DB.Migrator().CreateTable(&model.GroupUserModel{})
		if err != nil {
			log.Panicf("create table group_user err:%v", err)
		}
	}
	if !model.DB.Migrator().HasTable(&model.CollectModel{}) {
		err = model.DB.Migrator().CreateTable(&model.CollectModel{})
		if err != nil {
			log.Panicf("create table collection err:%v", err)
		}
	}
	if !model.DB.Migrator().HasTable(&model.UserTagModel{}) {
		err = model.DB.Migrator().CreateTable(&model.UserTagModel{})
		if err != nil {
			log.Panicf("create table user_tag err:%v", err)
		}
	}
	if !model.DB.Migrator().HasTable(&model.ReportModel{}) {
		err = model.DB.Migrator().CreateTable(&model.ReportModel{})
		if err != nil {
			log.Panicf("create table report err:%v", err)
		}
	}
	if !model.DB.Migrator().HasTable(&model.MessageModel{}) {
		err = model.DB.Migrator().CreateTable(&model.MessageModel{})
		if err != nil {
			log.Panicf("create table message err:%v", err)
		}
	}
	if !model.DB.Migrator().HasTable(&model.MomentModel{}) {
		err = model.DB.Migrator().CreateTable(&model.MomentModel{})
		if err != nil {
			log.Panicf("create table moment err:%v", err)
		}
	}
	if !model.DB.Migrator().HasTable(&model.MomentCommentModel{}) {
		err = model.DB.Migrator().CreateTable(&model.MomentCommentModel{})
		if err != nil {
			log.Panicf("create table moment comment err:%v", err)
		}
	}
	if !model.DB.Migrator().HasTable(&model.MomentLikeModel{}) {
		err = model.DB.Migrator().CreateTable(&model.MomentLikeModel{})
		if err != nil {
			log.Panicf("create table moment like err:%v", err)
		}
	}
	if !model.DB.Migrator().HasTable(&model.MomentTimelineModel{}) {
		err = model.DB.Migrator().CreateTable(&model.MomentTimelineModel{})
		if err != nil {
			log.Panicf("create table moment timeline err:%v", err)
		}
	}
	if !model.DB.Migrator().HasTable(&model.EmoticonModel{}) {
		err = model.DB.Migrator().CreateTable(&model.EmoticonModel{})
		if err != nil {
			log.Panicf("create table emoticon err:%v", err)
		}
	}
}

// 回滚数据库迁移
func down() {
	var err error
	if model.DB.Migrator().HasTable(&model.UserModel{}) {
		err = model.DB.Migrator().DropTable(&model.UserModel{})
		if err != nil {
			log.Panicf("drop table user err:%v", err)
		}
	}
	if model.DB.Migrator().HasTable(&model.ApplyModel{}) {
		err = model.DB.Migrator().DropTable(&model.ApplyModel{})
		if err != nil {
			log.Panicf("drop table applu err:%v", err)
		}
	}
	if model.DB.Migrator().HasTable(&model.GroupModel{}) {
		err = model.DB.Migrator().DropTable(&model.GroupModel{})
		if err != nil {
			log.Panicf("drop table group err:%v", err)
		}
	}
	if model.DB.Migrator().HasTable(&model.GroupUserModel{}) {
		err = model.DB.Migrator().DropTable(&model.GroupUserModel{})
		if err != nil {
			log.Panicf("drop table group_user err:%v", err)
		}
	}
	if model.DB.Migrator().HasTable(&model.CollectModel{}) {
		err = model.DB.Migrator().DropTable(&model.CollectModel{})
		if err != nil {
			log.Panicf("drop table collection err:%v", err)
		}
	}
	if model.DB.Migrator().HasTable(&model.FriendModel{}) {
		err = model.DB.Migrator().DropTable(&model.FriendModel{})
		if err != nil {
			log.Panicf("drop table friend err:%v", err)
		}
	}
	if model.DB.Migrator().HasTable(&model.UserTagModel{}) {
		err = model.DB.Migrator().DropTable(&model.UserTagModel{})
		if err != nil {
			log.Panicf("drop table user_tag err:%v", err)
		}
	}
	if model.DB.Migrator().HasTable(&model.ReportModel{}) {
		err = model.DB.Migrator().DropTable(&model.ReportModel{})
		if err != nil {
			log.Panicf("drop table report err:%v", err)
		}
	}
	if model.DB.Migrator().HasTable(&model.MessageModel{}) {
		err = model.DB.Migrator().DropTable(&model.MessageModel{})
		if err != nil {
			log.Panicf("drop table message err:%v", err)
		}
	}
	if model.DB.Migrator().HasTable(&model.MomentModel{}) {
		err = model.DB.Migrator().DropTable(&model.MomentModel{})
		if err != nil {
			log.Panicf("drop table moment err:%v", err)
		}
	}
	if model.DB.Migrator().HasTable(&model.MomentCommentModel{}) {
		err = model.DB.Migrator().DropTable(&model.MomentCommentModel{})
		if err != nil {
			log.Panicf("drop table moment comment err:%v", err)
		}
	}
	if model.DB.Migrator().HasTable(&model.MomentLikeModel{}) {
		err = model.DB.Migrator().DropTable(&model.MomentLikeModel{})
		if err != nil {
			log.Panicf("drop table moment like err:%v", err)
		}
	}
	if model.DB.Migrator().HasTable(&model.MomentTimelineModel{}) {
		err = model.DB.Migrator().DropTable(&model.MomentTimelineModel{})
		if err != nil {
			log.Panicf("drop table moment timeline err:%v", err)
		}
	}
	if !model.DB.Migrator().HasTable(&model.EmoticonModel{}) {
		err = model.DB.Migrator().DropTable(&model.EmoticonModel{})
		if err != nil {
			log.Panicf("drop table emoticon err:%v", err)
		}
	}
}
