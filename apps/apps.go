package apps

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/zuoyangs/go-grafana-dashboard/config"
)

func Run() {
	config.Init()

	//2. 初始化控制
	// 收到传递依赖关系：收到管理对象依赖
	// 2.1 user controller
	// userServiceImpl := userImpl.NewUserServiceImpl()
	// 2.2 token controller
	// tokenServiceImpl := tokenImpl.NewTokenServiceImpl(userServiceImpl)

	//  通过Ioc来完成依赖的装载, 完成了依赖的倒置(ioc 依赖对象注册)
	if err := ioc.Controller().Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 初始化Api Handler
	if err := ioc.ApiHandler().Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 跑在后台的http server
	httpServer := protocol.NewHttpServer()
	go func() {
		if err := httpServer.Run(); err != nil {
			fmt.Printf("start http server error, %s\n", err)
		}
	}()
	//3. 启动http协议服务器, 注册 handler路由
	// r := gin.Default()
	// ioc.ApiHandler().RouteRegistry(r.Group("/api/vblog"))

	// // 启动协议服务器
	// addr := conf.C().App.HttpAddr()
	// fmt.Printf("HTTP API监听地址: %s", addr)
	// err = r.Run(addr)

	// fmt.Println(err)
	// fmt.Println("清理工作")

	// 处理信号量
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

	// 等待信号的到来
	sn := <-ch
	fmt.Println(sn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	httpServer.Close(ctx)

	//
	fmt.Println("清理工作")

	log.Println("Start listen port 8080...")
	http.HandleFunc("/health", apps.healthCheck)
	http.HandleFunc("/serviceGet", apps.ServiceGet)
	http.ListenAndServe("0.0.0.0:8888", nil)

}
