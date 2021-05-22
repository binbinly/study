package cmd

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var cfg string

var rootCmd = &cobra.Command{
	Use:   "chat",
	Short: "chat app",
	Long:  `仿微信的即时通讯学习项目`,
	Args:  args,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("chat app start")
	},
}

func args(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("至少需要一个参数!")
	}
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(logicCmd)
	rootCmd.AddCommand(taskCmd)
	rootCmd.AddCommand(connectCmd)
}

func initConfig() {

}
