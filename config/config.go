package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"log"
	"os"
)

// init функция, которая закружает переменные окружения из .env
func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

// Конфиг для сервера и базы данных
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBPort   string
	DBName   string
}

// Конструктор конфига, который работает с переменными окружения из .env
func NewConfig() *Config {
	return &Config{
		Host:     getEnv("HOST"),
		Port:     getEnv("PORT"),
		User:     getEnv("USER"),
		Password: getEnv("PASSWORD"),
		DBPort:   getEnv("DB_PORT"),
		DBName:   getEnv("DB_NAME"),
	}
}

// Метод конфига. Делает Up миграцию из db/migrations
func (c *Config) MakeMigrations() {
	src := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		c.User, c.Password, c.Host, c.DBPort, c.DBName)

	db, err := sql.Open("postgres", src)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(db, "db/migrations"); err != nil {
		log.Fatal(err)
	}
}

// Функция находит переменную окружения по ключу. Вызывает панику, если переменная не найдена
func getEnv(key string) string {
	if envValue, ok := os.LookupEnv(key); !ok {
		panic(fmt.Sprintf("%s not found in env", key))
	} else {
		return envValue
	}
}
