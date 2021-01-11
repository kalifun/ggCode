package cmd

import (
	"fmt"
	"path"

	"github.com/kalifun/ggCode/pkg/folder"

	"github.com/spf13/cobra"
)

var src string
var target string
var project string

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
			isTrue, err := folder.PathExists(src)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(isTrue)
			isTrue, err = folder.PathExists(target)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(isTrue)
		}
	},
}

func Tips() {
	fmt.Println("ggCode gen --project=github.com/xxx/xxxx --src=/home/xxx/xxx.proto --target=/home/xxx/xxx")
}

func init() {
	Genrate.Flags().StringVarP(&project, "project", "p", "", "github.com/xxx/xxxx")
	Genrate.Flags().StringVarP(&src, "src", "s", "", "Source directory to read from")
	Genrate.Flags().StringVarP(&target, "target", "t", "", "Object files for generating code")
}
