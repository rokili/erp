package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Aa123456"
	dbname   = "erp_system"
)

func main() {
	// 连接到PostgreSQL服务器
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
		host, port, user, password)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL server:", err)
	}
	defer db.Close()

	// 创建数据库
	_, err = db.Exec("CREATE DATABASE " + dbname)
	if err != nil {
		// 如果数据库已存在，忽略错误
		if !strings.Contains(err.Error(), "already exists") {
			log.Fatal("Failed to create database:", err)
		}
		fmt.Println("Database already exists")
	} else {
		fmt.Println("Database created successfully")
	}

	// 连接到新创建的数据库
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	erpDB, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal("Failed to connect to ERP database:", err)
	}
	defer erpDB.Close()

	// 读取并执行SQL脚本
	sqlFile := "../init.sql"
	if len(os.Args) > 1 {
		sqlFile = os.Args[1]
	}

	sqlBytes, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Fatal("Failed to read SQL file:", err)
	}

	// 分割并执行SQL语句
	sqlStatements := strings.Split(string(sqlBytes), ";")
	for _, stmt := range sqlStatements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}

		_, err := erpDB.Exec(stmt)
		if err != nil {
			log.Printf("Failed to execute statement: %s\nError: %v", stmt, err)
			// 继续执行其他语句
		} else {
			fmt.Printf("Executed: %s\n", stmt[:min(len(stmt), 50)]+"...")
		}
	}

	fmt.Println("Database initialization completed")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
