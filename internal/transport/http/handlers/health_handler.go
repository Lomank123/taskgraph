package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Investigate fmt.Fprintf
	log.Println("Healthcheck hit!")
	fmt.Fprintf(w, "Hi there, %s", r.URL.Path[1:])
}
