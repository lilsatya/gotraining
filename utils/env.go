package utils

import (
	"os"
)

func getEnv(key, devEnv string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return devEnv
}
