package cmd

import (
	"fmt"
	"log"
	"mall/cmd/api"
	"mall/cmd/config"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"mall/cmd/migrate"
	"mall/cmd/version"
)

var rootCmd = &cobra.Command{
	Use:   "mall",
	Short: "mall app",
	Long:  `商城学习项目`,
	Args:  args,
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatalf("unrecognized command cmd: %v args: %v", cmd.Name(), args)
	},
}

func args(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("please select option")
	}
	return nil
}

//Execute 命令执行入口
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(config.StartCmd)
}
