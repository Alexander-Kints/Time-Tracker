package db

import (
	"TimeTracker/config"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// Устанавливает соединение с базой данных из PostgreSQL
func NewPostgresDB(cfg *config.Config) *sqlx.DB {
	src := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.DBPort, cfg.DBName)
	db, err := sqlx.Open("postgres", src)
	if err != nil {
		panic(err)
	}
	return db
}
