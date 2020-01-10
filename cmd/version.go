package cmd

import(
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version number of itool.",
	Run: func(cmd *cobra.Command, Args []string) {
		fmt.Println("The version number of itool is v1.0.0")
	},
}