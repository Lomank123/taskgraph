package utils

import (
	"fmt"
	"os"
)

// GetEnv returns the value of the environment variable specified by key.
// If the variable is not set, it prints a message and returns the provided fallback value.
func GetEnv(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		fmt.Println("Environment variable ", key, " not found, using fallback value: ", fallback)
		return fallback
	}
	return val
}