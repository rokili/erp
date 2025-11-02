package main

import (
	"erp/internal/config"
	"erp/internal/server"
	"log"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化服务器
	srv := server.New(cfg)

	// 启动服务器
	log.Printf("Starting server on %s:%s", cfg.Server.Host, cfg.Server.Port)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
