package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

func ENV(key string) string {
	err := godotenv.Load()
	if err != nil {
		LogError(err.Error())
	}

	return os.Getenv(key)
}
