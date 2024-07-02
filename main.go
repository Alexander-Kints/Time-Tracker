package main

import (
	"TimeTracker/config"
	"TimeTracker/db"
	"TimeTracker/handler"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	cfg := config.NewConfig()
	cfg.MakeMigrations()

	postgresDB := db.NewPostgresDB(cfg)
	defer postgresDB.Close()

	// HTML документация swagger
	http.HandleFunc("/docs", handler.DocsHandler)

	http.HandleFunc("/user/create", handler.CreateUserHandler(postgresDB))
	http.HandleFunc("/user/get/all", handler.GetAllUsersHandler(postgresDB))
	http.HandleFunc("/user/get/by/id", handler.GetUserDyIDHandler(postgresDB))
	http.HandleFunc("/user/get/by/name", handler.GetUsersByFilterHandler(postgresDB, "name"))
	http.HandleFunc("/user/get/by/surname", handler.GetUsersByFilterHandler(postgresDB, "surname"))
	http.HandleFunc("/user/get/by/patronymic", handler.GetUsersByFilterHandler(postgresDB, "patronymic"))
	http.HandleFunc("/user/get/by/address", handler.GetUsersByFilterHandler(postgresDB, "address"))
	http.HandleFunc("/user/get/statistics", handler.StatisticsHandler(postgresDB))
	http.HandleFunc("/user/delete", handler.DeleteUserHandler(postgresDB))
	http.HandleFunc("/user/update", handler.UpdateUserHandler(postgresDB))

	http.HandleFunc("/task/start", handler.StartTaskHandler(postgresDB))
	http.HandleFunc("/task/finish", handler.FinishTaskHandler(postgresDB))

	// эндпоинт для имитации запроса к стороннему сервису
	http.HandleFunc("/info", handler.InfoHandler(postgresDB))

	log.Println("server start at port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		panic(err)
	}
}
