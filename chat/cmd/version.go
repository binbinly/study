package cmd

import (
	"fmt"
	"os"

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

		fmt.Fprint(os.Stdout, "chat version ", output)
	},
}