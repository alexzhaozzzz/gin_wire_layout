package main

import (
	"github.com/alexzhaozzzz/gin_wire_layout/internal/model/mysql"
	"github.com/alexzhaozzzz/gin_wire_layout/pkg/bootstrap"
	"github.com/alexzhaozzzz/gin_wire_layout/pkg/serverx"
	"github.com/spf13/pflag"
)

// 主程序入口
func main() {
	// 定义一个命令行Flag来指定配置文件路径
	path := pflag.StringP("config", "c", "", "Path to the config file")
	pflag.Parse()

	// 加载配置文件
	err := bootstrap.LoadConfig(*path)
	if err != nil {
		panic("load config error: " + err.Error())
	}
	// 初始化日志系统
	bootstrap.InitLogger(bootstrap.WithOption("appName", "aa"))
	// 确保程序退出时日志能同步
	defer bootstrap.Sync()

	// 创建数据库连接
	ds := mysql.NewDefaultMysql()
	// 创建HTTP服务器实例
	srv := serverx.NewAppServer()
	// 注册服务器关闭时的处理函数，确保数据库连接能被正确关闭
	srv.RegisterOnShutdown(func() {
		if ds != nil {
			ds.Close()
		}
	})
	//初始化路由
	router := initRouter(ds)
	// 启动HTTP服务器，应用中间件和路由
	srv.Run(router)
}
