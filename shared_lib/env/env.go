package env

import (
	"os"
)

func GetLoggerAddress() string {
	host := getEnv("LOGGER_HOST", "logger.uni-ulm.de")
	port := getEnv("LOGGER_PORT", "4432")
	return host + ":" + port
}

// Get env var or default
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
