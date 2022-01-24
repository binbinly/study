package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"chat-micro/cmd/connect"
	"chat-micro/cmd/logic"
	"chat-micro/cmd/migrate"
	"chat-micro/cmd/seed"
	"chat-micro/cmd/task"
	"chat-micro/cmd/version"
)

var rootCmd = &cobra.Command{
	Use:   "chat-micro",
	Short: "chat-micro app",
	Long:  `仿微信的即时通讯学习项目`,
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
	rootCmd.AddCommand(migrate.LogicCmd)
	rootCmd.AddCommand(logic.StartCmd)
	rootCmd.AddCommand(task.StartCmd)
	rootCmd.AddCommand(connect.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(seed.StartCmd)
}