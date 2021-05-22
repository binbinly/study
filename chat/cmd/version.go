package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of chat",
	Long:  `All software has versions. This is chat`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chat Static Site Generator v0.1 -- HEAD")
	},
}