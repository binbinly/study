package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/cobra"

	"chat/app/logic/conf"
	"chat/app/logic/model"
	"chat/pkg/crypt/auth"
	"chat/pkg/log"
)

var clear bool
var n int

func init() {
	seedCmd.Flags().StringVarP(&cfg, "config", "c", "", "config file (default is $ROOT/config/logic.yaml)")
	seedCmd.Flags().BoolVar(&clear, "clear", false, "clear db table data")
	seedCmd.Flags().IntVarP(&n, "number", "n", 100000, "set data number")
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "chat seed data",
	Run: func(cmd *cobra.Command, args []string) {
		if cfg == "" {
			cfg = "./config/logic.yaml"
		}
		Init()
		Seed()
	},
}

func Init() {
	conf.Init(cfg)
	model.Init(&conf.Conf.MySQL)
	log.InitLog(log.NewConfig())
}

func Seed() {
	if clear {
		clearData()
	} else {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			seedUser()
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			seedFriend()
		}()
		wg.Wait()
	}
}

func seedUser() {
	fmt.Println("seed user")

	password, _ := auth.Encrypt("123456")
	us := make([]*model.UserModel, n)
	for i := 0; i < n; i++ {
		u := &model.UserModel{
			Username: fmt.Sprintf("test%d", i),
			Password: password,
			Phone:    int64(15888888888 + i),
		}
		us[i] = u
	}
	var end int
	batchSize := 1000 //批处理大小
	// 按大小进行批处理插入，防好友太多，插入数据库失败
	for i := 0; i < n; i += batchSize {
		end = i + batchSize
		if end > n {
			end = n
		}
		sub := us[i:end]
		now := time.Now().Unix()
		sql := "INSERT INTO `user` (`username`,`nickname`,`password`,`phone`,`email`,`avatar`,`gender`,`status`,`sign`,`area`,`created_at`,`updated_at`) VALUES "
		for _, um := range sub {
			sql += fmt.Sprintf("('%v','','%v',%v,'','',1,0,'','',%v,%v),", um.Username, um.Password, um.Phone, now, now)
		}
		sql = sql[0 : len(sql)-1]
		fmt.Println("save", i)
		model.DB.Exec(sql)
	}
}

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
	model.DB.Model(&model.FriendModel{}).CreateInBatches(fs, 1000)
}

//清空表数据
func clearData() {
	model.DB.Exec("truncate user")
	model.DB.Exec("truncate friend")
	model.DB.Exec("truncate moment")
	model.DB.Exec("truncate moment_timeline")
}
