package cmd

import (
	"fmt"

	"github.com/kalifun/ggCode/gengrpc"

	"github.com/kalifun/ggCode/common"
	"github.com/kalifun/ggCode/pkg/font"

	"github.com/kalifun/ggCode/config"

	"github.com/spf13/cobra"
)

// var src string
// var target string
// var project string

var Genrate = &cobra.Command{
	Use:   "gen",
	Short: "read proto file,then genrate grpc code",
	Long:  "-src [path] -target [path]",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		config.InitConfig()
		if config.ConfSvr.Genrate.Src == "" || config.ConfSvr.Genrate.Target == "" || config.ConfSvr.Genrate.Project == "" || config.ConfSvr.Genrate.GrpcWorkSpace == "" {
			msg := font.NewColorFont(common.Empty)
			fmt.Println(msg.Red())
			return
		} else {
			gengrpc.GenGrpcMain()
		}
	},
}

func Tips() {
	fmt.Println("ggCode gen --project=github.com/xxx/xxxx --src=/home/xxx/xxx.proto --target=/home/xxx/xxx")
}

// func init() {
// 	Genrate.Flags().StringVarP(&project, "project", "p", "", "github.com/xxx/xxxx")
// 	Genrate.Flags().StringVarP(&src, "src", "s", "", "Source directory to read from")
// 	Genrate.Flags().StringVarP(&target, "target", "t", "", "Object files for generating code")
// }
