package main

import (
	"fmt"
	"os"

	"github.com/kalifun/ggCode/cmd"
)

func init() {
	cmd.StartCmd.AddCommand(cmd.Version)
	cmd.StartCmd.AddCommand(cmd.Genrate)
}

func main() {
	if err := cmd.StartCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
