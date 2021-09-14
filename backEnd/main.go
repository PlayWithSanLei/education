package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/impact-eintr/education/global"
	"github.com/impact-eintr/education/internal/router"
	"go.uber.org/zap"
)

func init() {
	// 加载配置文件
	global.Init()
}

// @title EDU后台管理系统
// @version 1.0.0
// @description 华迪的作业
func main() {
	r, err := router.NewRouter()
	if err != nil {
		return
	}

	server := &http.Server{
		Addr:    net.JoinHostPort("127.0.0.1", "8080"),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("Listen: ", zap.Error(err))
			panic(err)
		}
	}()

	// 优雅关机
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// TODO 关闭数据库链接
	if err := server.Shutdown(ctx); err != nil {
		zap.L().Error("shutdown", zap.Error(err))
	}

	zap.L().Info("Server exit")

}
