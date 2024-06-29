package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	godotenv.Load(".env")

	val := os.Getenv(key)
	if val == "" {
		log.Fatal(key, " is not found in the environment")
	}
	return val
}
