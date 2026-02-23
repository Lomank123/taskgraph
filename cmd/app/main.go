package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"taskgraph/config"
	"taskgraph/internal/services"
	httpTransport "taskgraph/internal/transport/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// Swagger docs
	_ "taskgraph/docs"
)

// @title taskgraph API
// @version 1.0
// @description This is a task management server.
// @BasePath /api
func main() {
	log.Println("Loading config...")
	cfg := config.LoadConfig()

	log.Println("Connecting to database...")
	db, err := gorm.Open(postgres.Open(cfg.Postgres.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	log.Println("Database connected!")
	
	taskService := services.NewTaskService(db)
	router := httpTransport.NewRouter(taskService)
	server := httpTransport.NewHTTPServer(router)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Starting server on port ", cfg.App.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("FATAL ERROR: ", err)
		}
	}()

	<-quit
	log.Println("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// Closing database connections
	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.Close()
		log.Println("Database connection closed successfully")
	}

	log.Println("Server exited properly")
}
