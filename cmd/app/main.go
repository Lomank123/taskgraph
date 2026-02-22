package main

import (
	"log"
	"taskgraph/config"
	"taskgraph/internal/services"
	httpTransport "taskgraph/internal/transport/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.Println("Loading config...")
	cfg := config.LoadConfig()

	log.Println("Connecting to database...")
	db, err := gorm.Open(postgres.Open(cfg.Postgres.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	log.Println("Database connected!")

	log.Println("Starting server on port ", cfg.App.Port)
	taskService := services.NewTaskService(db)
	router := httpTransport.NewRouter(taskService)
	server := httpTransport.NewHTTPServer(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("FATAL ERROR: ", err)
	}
}
