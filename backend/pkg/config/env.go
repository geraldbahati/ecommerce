package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	DbUrl string
}

func LoadConfig() Config {
	if err := godotenv.Load(".env"); err != nil {
		// log.Fatal("Error loading .env file: ", err)

		// Added this for testing and debugging reasons
		// To be removed
		log.Println("Error loadding .env file:", err)
		return Config{
			Port: "8000",
			DbUrl: "postgresql://postgres:staphone@16@localhost:5432/staphone?sslmode=disable",
		}
	}

	return Config{
		Port:  getEnv("PORT", "8080"),
		DbUrl: getEnv("DB_URL", ""),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
