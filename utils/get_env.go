package utils

import (
	"log"
	"os"
)

// GetEnv returns the value of the environment variable specified by key.
// If the variable is not set, it prints a message and returns the provided fallback value.
func GetEnv(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("Environment variable %s not found, using fallback value: %s", key, fallback)
		return fallback
	}
	return val
}
