package cmd

import (
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "ggCode",
	Short: "generate grpc code",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if len(args) < 1 {
			cmd.Help()
			return
		}

	},
}
