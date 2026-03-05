package env

import (
	"os"
	"strconv"
)

// Retrieves string environment variable
// key = variable name, defaultValue = value to use if not found
func GetEnvString(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value // returns the value from .env file
	}
	return defaultValue // otherwise return default value
}

// Retrieves integer environment variable
func GetEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
