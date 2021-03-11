package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config : get env value
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error Loading .env File")
	}
	return os.Getenv(key)
}
