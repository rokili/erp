package server

import (
	"erp/internal/config"
	"erp/internal/handler"
	"erp/internal/repository"
	"erp/internal/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Server struct {
	config *config.Config
	router *gin.Engine
	db     *sqlx.DB
}

func New(cfg *config.Config) *Server {
	// 初始化数据库连接
	db, err := sqlx.Connect("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 初始化Gin路由器
	router := gin.Default()

	// 初始化依赖
	repo := repository.New(db)
	svc := service.New(repo)
	handler := handler.New(svc)

	// 注册路由
	registerRoutes(router, handler)

	return &Server{
		config: cfg,
		router: router,
		db:     db,
	}
}

func registerRoutes(router *gin.Engine, handler *handler.Handler) {
	// 提供静态文件服务
	router.Static("/web", "./web")

	// 默认页面重定向到 web 界面
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/web")
	})

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 财务相关路由
	finance := router.Group("/api/finance")
	{
		finance.POST("/vouchers", handler.CreateVoucher)
		finance.GET("/vouchers/:id", handler.GetVoucher)
		finance.GET("/vouchers", handler.ListVouchers)
		finance.POST("/accounts", handler.CreateAccount)
		finance.GET("/accounts/:code", handler.GetAccount)
		finance.GET("/accounts", handler.ListAccounts)
	}

	// 采购相关路由
	purchase := router.Group("/api/purchase")
	{
		purchase.POST("/orders", handler.CreatePurchaseOrder)
		purchase.GET("/orders/:id", handler.GetPurchaseOrder)
		purchase.GET("/orders", handler.ListPurchaseOrders)
		purchase.PUT("/orders/:id/approve", handler.ApprovePurchaseOrder)
		purchase.PUT("/orders/:id/close", handler.ClosePurchaseOrder)
	}

	// 销售相关路由
	sales := router.Group("/api/sales")
	{
		sales.POST("/orders", handler.CreateSalesOrder)
		sales.GET("/orders/:id", handler.GetSalesOrder)
		sales.GET("/orders", handler.ListSalesOrders)
		sales.PUT("/orders/:id/approve", handler.ApproveSalesOrder)
		sales.PUT("/orders/:id/close", handler.CloseSalesOrder)
	}
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%s", s.config.Server.Host, s.config.Server.Port)
	return s.router.Run(addr)
}

func (s *Server) Close() error {
	return s.db.Close()
}
