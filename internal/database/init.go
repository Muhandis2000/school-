// internal/database/init.go
package database

import (
	"Final_project/config"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

// ConnectDB устанавливает соединение с PostgreSQL
func ConnectDB() {
	cfg := config.Cfg
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.SSLMode)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	DB = db
}
