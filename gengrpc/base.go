package gengrpc

import (
	"fmt"
	"os"
	"path"

	"github.com/kalifun/ggCode/pkg/folder"

	"github.com/kalifun/ggCode/config"
)

func GenGrpcMain() {
	if path.IsAbs(config.ConfSvr.Genrate.Src) == false {
		fmt.Println("SRC is not an absolute path")
		return
	}
	if path.IsAbs(config.ConfSvr.Genrate.Target) == false {
		fmt.Println("Target is not an absolute path")
		return
	}

	// 判断文件是否存在
	ok, err := folder.PathExists(config.ConfSvr.Genrate.Src)
	if err != nil {
		fmt.Println(err)
		return
	}
	if ok == false {
		fmt.Println("File does not exist")
		return
	}
	ok, err = folder.PathExists(config.ConfSvr.Genrate.Target)
	if err != nil {
		fmt.Println(err)
		return
	}
	if ok == false {
		err := os.Mkdir(config.ConfSvr.Genrate.Target, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

type GenGrpcType struct {
	SrcPath    string
	TargetPath string
}

const mainFile = `package main

import (
	"fmt"
	pb "git.code.oa.com/annotate-group/grpc/proto/statistics_config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"net"
)

func init() {
	sqlSet := global.ConfigServer.Mysql
	//user=postgres password=tom dbname=test host=127.0.0.1 port=5432 sslmode=disable
	/**
	 * MySQL 配置
	 * 注册驱动
	 * orm.RegisterDriver("mysql", orm.DR_MySQL)
	 * mysql用户：root ，root的秘密：tom ， 数据库名称：test ， 数据库别名：default
	 * orm.RegisterDataBase("default", "mysql", "root:tom@/test?charset=utf8")
	 */
	sqlConn := sqlSet.Username + ":" + sqlSet.Password + "@(" + sqlSet.Path + ")/" + sqlSet.Dbname + "?" + sqlSet.Config
	global.Oplog.Info("Start Init Mysql")
	global.Oplog.Info(sqlConn)
	RegisterDriver := orm.RegisterDriver("mysql", orm.DRMySQL)
	if RegisterDriver != nil {
		global.Oplog.Error(RegisterDriver)
	}
	RegisterDataBase := orm.RegisterDataBase("default",
		"mysql", sqlConn, sqlSet.MaxIdleConns, sqlSet.MaxOpenConns)
	if RegisterDataBase != nil {
		global.Oplog.Error(RegisterDataBase)
	}
	orm.Debug = true
	orm.RegisterModel() // 自动new表
	RunSyncDb := orm.RunSyncdb("default", false, true)
	if RunSyncDb != nil {
		global.Oplog.Error(RunSyncDb)
	}
	//utils.ClearLock("test")
}

func main() {
	server := grpc.NewServer()
	pb.RegisterManpowerStatisticsServer(server, &controller.StatisticsServer{})
	if global.ConfigServer.System.RPCPORT == "" {
		global.Oplog.Error("Can't Found GRPC PORT")
	}
	port := global.ConfigServer.System.RPCPORT
	fmt.Printf("address is :%s\n", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		global.Oplog.Error(err)
	}
	global.Oplog.Info("Service has started, start port ", global.ConfigServer.System.RPCPORT)
	if err := server.Serve(lis); err != nil {
		global.Oplog.Error(err)
	}

}
`
