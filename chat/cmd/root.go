package cmd

import (
	"fmt"
	"os"
	"os/exec"

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
		Error(cmd, args, errors.New("unrecognized command"))
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
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(logicCmd)
	rootCmd.AddCommand(taskCmd)
	rootCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(seedCmd)
}

func initConfig() {

}

//ExecuteCommand 执行命令
func ExecuteCommand(name string, subName string, args ...string) (string, error) {
	args = append([]string{subName}, args...)

	cmd := exec.Command(name, args...)
	bytes, err := cmd.CombinedOutput()

	return string(bytes), err
}

//Error 输出错误
func Error(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "execute %s args:%v error:%v\n", cmd.Name(), args, err)
	os.Exit(1)
}
