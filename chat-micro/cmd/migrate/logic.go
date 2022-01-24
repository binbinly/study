package migrate

import (
	"fmt"

	"github.com/spf13/cobra"

	"chat-micro/app/logic/conf"
	"chat-micro/cmd/migrate/migration/logic"
	"chat-micro/internal/orm"
)

var (
	LogicCmd = &cobra.Command{
		Use:     "migrate",
		Short:   "Initialize the database",
		Example: "mall migrate -c config/logic/default.yml",
		Run: func(cmd *cobra.Command, args []string) {
			startLogic()
		},
	}
)

func init() {
	LogicCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/logic/default.yml", "Start server with provided configuration file")
	LogicCmd.PersistentFlags().BoolVarP(&generate, "generate", "g", false, "generate migration file")
}

func startLogic() {

	if !generate {
		//1. 加载迁移
		logic.Init()
		//2. 读取配置
		conf.Init(configYml)
		//3. 初始化数据库链接
		orm.Init(&conf.Conf.MySQL)
		//4. 执行迁移
		up()
	} else {
		fmt.Println(`generate logic migration file`)
		_ = genFile("logic")
	}
}
