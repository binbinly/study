package config

import (
	"fmt"
	"github.com/spf13/cobra"

	"mall/app/conf"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:     "config",
		Short:   "Get Application config info",
		Example: "mall config -c config/default.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/default.yml", "Start server with provided configuration file")
}

func run() {
	conf.Init(configYml)
	fmt.Printf("config:%+v\n", conf.Conf)
}
