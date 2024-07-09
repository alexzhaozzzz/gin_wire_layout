// Created on 2024-7-8 18:10:25
// @author alex
// email 1368751885@qq.com
// 通用server启动类，完成配置加载，缓存初始化，数据库初始化，启动参数绑定

package serverx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/alexzhaozzzz/gin_wire_layout/pkg/bootstrap"
	"github.com/alexzhaozzzz/gin_wire_layout/pkg/colorx"
	"github.com/alexzhaozzzz/gin_wire_layout/pkg/util"
)

// AppServer 代表当前服务端实例
type AppServer struct {
	f func()
}

// NewAppServer 创建server实例
func NewAppServer() *AppServer {
	return &AppServer{}
}

// IRouter 加载路由，使用侧提供接口，实现侧需要实现该接口
type IRouter interface {
	Load(engine *gin.Engine)
}

// Run server的启动入口
// 加载路由, 启动服务
func (s *AppServer) Run(rs ...IRouter) {
	httpHost := viper.GetString("server.http.host")
	httpPort := viper.GetInt("server.http.port")
	if httpHost == "" || httpPort <= 0 {
		panic("get http config err")
	}

	var wg sync.WaitGroup
	wg.Add(1)
	// 设置gin启动模式，必须在创建gin实例之前
	if bootstrap.IsDevelopment() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	g := gin.New()

	s.routerLoad(g, rs...)

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%d", httpHost, httpPort),
		Handler: g,
	}
	if s.f != nil {
		srv.RegisterOnShutdown(s.f)
	}

	// graceful shutdown
	sgn := make(chan os.Signal, 1)
	signal.Notify(sgn, syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGQUIT)

	go func() {
		<-sgn
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			//TODO: 增加记录
		}
		wg.Done()
	}()

	local := util.GetLocalIP()
	fmt.Println(colorx.Green("\r\nServer run at:"))
	fmt.Printf("-  Local:   %s://localhost:%d/ \r\n", "http", httpPort)
	fmt.Printf("-  Network: %s://%s:%d/ \r\n", "http", local, httpPort)

	if bootstrap.IsDevelopment() {
		fmt.Println(colorx.Green("Swagger run at:"))
		fmt.Printf("-  Local:   http://localhost:%d/swagger/index.html \r\n", httpPort)
		fmt.Printf("-  Network: %s://%s:%d/swagger/index.html \r\n", "http", local, httpPort)
		fmt.Printf("\r\n %s Enter Control + C Shutdown Server \r\n", time.Now().Format(time.DateTime))
	}

	err := srv.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server Start Err %s \r\n", err.Error())
			return
		}
	}

	wg.Wait()

	fmt.Printf("%s Server Exited ... \r\n", time.Now().Format(time.DateTime))
}

// RouterLoad 加载自定义路由
func (s *AppServer) routerLoad(g *gin.Engine, rs ...IRouter) *AppServer {
	for _, r := range rs {
		r.Load(g)
	}
	return s
}

// RegisterOnShutdown 注册shutdown后的回调处理函数，用于清理资源
func (s *AppServer) RegisterOnShutdown(_f func()) {
	s.f = _f
}
