package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	DbUrl           string
	DefaultPageSize int32
	DefaultPage     int32
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		Port:            getEnv("PORT", "8080"),
		DbUrl:           getEnv("DB_URL", ""),
		DefaultPageSize: 100,
		DefaultPage:     1,
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
