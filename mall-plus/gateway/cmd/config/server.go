package config

import (
	"fmt"

	"github.com/spf13/cobra"

	"gateway/conf"
)

var (
	configYml string
	//StartCmd config cmd
	StartCmd = &cobra.Command{
		Use:     "config",
		Short:   "Get Application config info",
		Example: "mall config -c default.yaml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "default.yaml", "Start server with provided configuration file")
}

func run() {
	conf.Init(configYml)
	fmt.Printf("config:%+v\n", conf.Conf)
}
