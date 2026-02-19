package main

import (
	"fmt"
	"log"
	"net/http"
	"taskgraph/config"
	httpTransport "taskgraph/internal/transport/http"
)

func main() {
	log.Println("Loading config...")
	cfg := config.LoadConfig()

	log.Println("Starting server...")
	router := httpTransport.SetupRouter()

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", cfg.App.Port),
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("FATAL ERROR: ", err)
	}
}
