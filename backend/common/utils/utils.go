package utils

import (
	"os"
	"strconv"
)

func GetEnvString(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func GetEnvInt64(key string, defaultValue int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intValue
		}
	}
	return defaultValue
}
