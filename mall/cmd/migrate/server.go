package migrate

import (
	"bytes"
	"fmt"
	"strconv"
	"text/template"
	"time"

	"github.com/spf13/cobra"

	"mall/app/conf"
	"mall/app/model"
	"mall/cmd/migrate/migration"
	_ "mall/cmd/migrate/migration/version"
	"mall/pkg/utils"
)

var (
	configYml string
	generate  bool
	StartCmd  = &cobra.Command{
		Use:     "migrate",
		Short:   "Initialize the database",
		Example: "mall migrate -c config/default.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/default.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&generate, "generate", "g", false, "generate migration file")
}

func run() {

	if !generate {
		fmt.Println(`start init`)
		//1. 读取配置
		conf.Init(configYml)
		initDB()
	} else {
		fmt.Println(`generate migration file`)
		_ = genFile()
	}
}

func initDB() {
	//3. 初始化数据库链接
	model.Init(&conf.Conf.MySQL)
	//4. 数据库迁移
	fmt.Println("数据库迁移开始")
	_ = migrateModel()
	fmt.Println(`数据库基础数据初始化成功`)
}

func migrateModel() error {
	db := model.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")

	err := db.Debug().AutoMigrate(&model.Migration{})
	if err != nil {
		return err
	}
	migration.Migrate.SetDb(db.Debug())
	migration.Migrate.Migrate()
	return err
}

func genFile() error {
	t1, err := template.ParseFiles("template/migrate.template")
	if err != nil {
		return err
	}
	m := map[string]string{}
	m["GenerateTime"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	m["Package"] = "version"
	var b1 bytes.Buffer
	err = t1.Execute(&b1, m)
	if err != nil {
		return err
	}
	return utils.FileCreate(b1, "./cmd/migrate/migration/version/"+m["GenerateTime"]+"_migrate.go")
}
