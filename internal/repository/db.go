package repository

import (
	"erp/internal/config"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func InitDB() {
	cfg := config.Load()

	// 构建PostgreSQL连接字符串
	connStr := "host=" + cfg.Database.Host +
		" port=" + fmt.Sprintf("%d", cfg.Database.Port) +
		" user=" + cfg.Database.User +
		" password=" + cfg.Database.Password +
		" dbname=" + cfg.Database.Name +
		" sslmode=" + cfg.Database.SSLMode

	var err error
	db, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 测试连接
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Database connection established")
}

func GetDB() *sqlx.DB {
	return db
}
