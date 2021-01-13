package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/kalifun/ggCode/common"
	"github.com/spf13/viper"
)

var (
	ConfSvr Server
	Viper   *viper.Viper
)

type Server struct {
	Genrate Genrate `json:"genrate" yaml:"genrate"`
}

/*
Src Proto 绝对路径
Target 需要生成文件的绝对路径(自动创建目录)
Project 生成文件的git路径
ProtoWorkSpace grpc server路径
*/
type Genrate struct {
	Src           string `json:"src" yaml:"src"`
	Target        string `yaml:"target"`
	Project       string `yaml:"project"`
	GrpcWorkSpace string `yaml:"grpcworkspace"`
}

func InitConfig() {
	v := viper.New()
	v.SetConfigFile(common.ConfigPath)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&ConfSvr); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&ConfSvr); err != nil {
		fmt.Println(err)
	}
	Viper = v
}
