package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DBConn() (db *sql.DB) {
	dsn := "root:1@tcp(127.0.0.1:3306)/go_api"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil
	}

	// Kiểm tra kết nối
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
		return nil
	}

	fmt.Println("Successfully connected to MySQL!")

	return db
}
