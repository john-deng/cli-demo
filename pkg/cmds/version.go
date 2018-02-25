package cmds

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cli-demo",
	Long:  `All software has versions. This is cli-demo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cli-demo v0.1")
	},
}