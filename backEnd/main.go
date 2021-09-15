package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/impact-eintr/education/internal/dao/db"
	"github.com/impact-eintr/education/internal/router"
	"github.com/impact-eintr/education/pkg/logger"
	"github.com/impact-eintr/education/pkg/rbac"
	sf "github.com/impact-eintr/education/pkg/snowflake"
	"github.com/impact-eintr/education/pkg/trans"
	"go.uber.org/zap"
)

// 加载配置文件
func init() {
	// 初始化日志
	_ = logger.L
	zap.L().Debug("logger init success...")

	// 初始化ID生成器
	_ = sf.S
	zap.L().Debug("ID init success...")

	// 初始化翻译器设置
	_ = trans.T
	zap.L().Debug("Trans init success...")

	// 初始化sql设置
	_ = db.DB
	zap.L().Debug("DB init success...")

	// 初始化RBAC
	_ = rbac.R
	zap.L().Debug("RBAC init success...")

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
