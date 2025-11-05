package main

import (
	"erp/internal/config"
	"erp/internal/handler"
	"erp/internal/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	repository.InitDB()

	// 加载配置
	cfg := config.Load()

	// 设置Gin模式
	if cfg.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建Gin引擎
	r := gin.Default()

	// 提供静态文件服务
	r.Static("/web", "./web")

	// 注册路由
	handler.RegisterRoutes(r)

	// 启动服务器
	addr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
