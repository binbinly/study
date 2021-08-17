package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of chat",
	Long:  `All software has versions. This is chat`,
	Run: func(cmd *cobra.Command, args []string) {
		output, err := ExecuteCommand("git", "describe", "--tags")
		if err != nil {
			Error(cmd, args, err)
		}
		fmt.Fprintln(os.Stdout, "chat version ", output)
		fmt.Fprintln(os.Stdout, "go version ", runtime.Version())
		fmt.Fprintln(os.Stdout, "Compiler ", runtime.Compiler)
		fmt.Fprintln(os.Stdout, "Platform ", fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
	},
}
