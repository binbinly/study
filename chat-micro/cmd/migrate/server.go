package migrate

import (
	"bytes"
	"fmt"
	"strconv"
	"text/template"
	"time"

	"chat-micro/cmd/migrate/migration"
	"chat-micro/internal/orm"
	"chat-micro/pkg/util"
)

var (
	configYml string
	generate  bool
)

func up() {
	//4. 数据库迁移
	fmt.Println("数据库迁移开始")
	_ = migrateModel()
	fmt.Println(`数据库基础数据初始化成功`)
}

func migrateModel() error {
	db := orm.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")

	err := db.Debug().AutoMigrate(&orm.Migration{})
	if err != nil {
		return err
	}
	migration.Migrate.SetDb(db.Debug())
	migration.Migrate.Migrate()
	return err
}

func genFile(server string) error {
	t1, err := template.ParseFiles("template/migrate.template")
	if err != nil {
		return err
	}
	m := map[string]string{}
	m["GenerateTime"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	m["Package"] = server
	var b1 bytes.Buffer
	err = t1.Execute(&b1, m)
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("./cmd/migrate/migration/%s/%s_migrate.go", m["Package"], m["GenerateTime"])
	fmt.Println("filename", filename)
	return util.FileCreate(b1, filename)
}
