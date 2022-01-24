package seed

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"

	"chat-micro/app/logic/conf"
	"chat-micro/app/logic/model"
	"chat-micro/internal/orm"
)

var (
	configYml string
	emo       bool
	n         int
	clear     bool
	StartCmd  = &cobra.Command{
		Use:          "seed",
		Short:        "Seed data",
		Example:      "chat-micro seed -c config/logic/default.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/logic/default.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&emo, "emoticon", "e", false, "sync remote emoticon data")
	StartCmd.PersistentFlags().BoolVar(&clear, "clear", false, "Clear data")
	StartCmd.PersistentFlags().IntVarP(&n, "number", "n", 100, "Create data row number")
}

func setup() {
	// init conf
	conf.Init(configYml)
	// init db
	orm.Init(&conf.Conf.MySQL)
}

func run() {
	if emo {
		err := SyncBQB()
		if err != nil {
			fmt.Printf("err:%v\n", err)
		}
		fmt.Println("sync emoticon success")
		return
	}
	if clear {
		clearData()
	} else {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			//seedUser()
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			seedFriend()
		}()
		wg.Wait()
	}
}

//func seedUser() {
//	fmt.Println("seed user")
//
//	password, _ := auth.Encrypt("123456")
//	us := make([]*model.UserModel, n)
//	for i := 0; i < n; i++ {
//		u := &model.UserModel{
//			Username: fmt.Sprintf("test%d", i),
//			Password: password,
//			Phone:    int64(15888888888 + i),
//		}
//		us[i] = u
//	}
//	var end int
//	batchSize := 1000 //批处理大小
//	// 按大小进行批处理插入，防好友太多，插入数据库失败
//	for i := 0; i < n; i += batchSize {
//		end = i + batchSize
//		if end > n {
//			end = n
//		}
//		sub := us[i:end]
//		now := time.Now().Unix()
//		sql := "INSERT INTO `user` (`username`,`nickname`,`password`,`phone`,`email`,`avatar`,`gender`,`status`,`sign`,`area`,`created_at`,`updated_at`) VALUES "
//		for _, um := range sub {
//			sql += fmt.Sprintf("('%v','','%v',%v,'','',1,0,'','',%v,%v),", um.Username, um.Password, um.Phone, now, now)
//		}
//		sql = sql[0 : len(sql)-1]
//		fmt.Println("save", i)
//		orm.DB.Exec(sql)
//	}
//}

func seedFriend() {
	fmt.Println("seed friend")

	var fs []*model.FriendModel
	for i := 4; i < n/2; i++ {
		f1 := &model.FriendModel{
			UserID:   uint32(i),
			FriendID: uint32(n - i),
			LookMe:   1,
			LookHim:  1,
		}
		f2 := &model.FriendModel{
			UserID:   uint32(n - i),
			FriendID: uint32(i),
			LookMe:   1,
			LookHim:  1,
		}
		fs = append(fs, f1, f2)
		if i > 5000 {
			continue
		}
		f3 := &model.FriendModel{
			UserID:   1,
			FriendID: uint32(i),
			LookMe:   1,
			LookHim:  1,
		}
		f4 := &model.FriendModel{
			UserID:   2,
			FriendID: uint32(i),
			LookMe:   1,
			LookHim:  1,
		}
		f5 := &model.FriendModel{
			UserID:   3,
			FriendID: uint32(i),
			LookMe:   1,
			LookHim:  1,
		}
		fs = append(fs, f3, f4, f5)
	}
	fmt.Println("init friend finish")
	orm.DB.Model(&model.FriendModel{}).CreateInBatches(fs, 1000)
}

//清空表数据
func clearData() {
	orm.DB.Exec("truncate user")
	orm.DB.Exec("truncate friend")
	orm.DB.Exec("truncate moment")
	orm.DB.Exec("truncate moment_timeline")
}
