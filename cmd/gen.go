package cmd

import (
	"fmt"
	"path"

	"github.com/spf13/cobra"
)

var src string
var target string

var Genrate = &cobra.Command{
	Use:   "gen",
	Short: "read proto file,then genrate grpc code",
	Long:  "-src [path] -target [path]",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if src == "" || target == "" {
			Tips()
		} else {
			if path.IsAbs(src) == false {
				fmt.Println("SRC is not an absolute path")
			}

			if path.IsAbs(target) == false {
				fmt.Println("TARGET is not an absolute path")
			}
		}
	},
}

func Tips() {
	fmt.Println("ggCode gen --src=/home/xxx/xxx.proto --target=/home/xxx/xxx")
}

func init() {
	Genrate.Flags().StringVarP(&src, "src", "s", "", "Source directory to read from")
	Genrate.Flags().StringVarP(&target, "target", "t", "", "Object files for generating code")
}
