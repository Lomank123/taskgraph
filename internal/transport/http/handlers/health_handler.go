package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Healthcheck hit!")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
