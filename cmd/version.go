package cmd

import (
	"fmt"

	"github.com/kalifun/ggCode/common"

	"github.com/spf13/cobra"
)

var Version = &cobra.Command{
	Use:   "version",
	Short: "print version",
	Long:  "Version",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println(common.Version)
	},
}
