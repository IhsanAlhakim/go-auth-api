package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBName     string
	DBUsername string
	DBPassword string

	SessionID string
	EcryptKey string

	RedisAddr     string
	RedisUsername string
	RedisPassword string
	RedisDB       string

	Port string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		DBName:        os.Getenv("DB_NAME"),
		DBUsername:    os.Getenv("DB_USERNAME"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		SessionID:     os.Getenv("SESSION_ID"),
		EcryptKey:     os.Getenv("ENCRYPTION_KEY"),
		RedisAddr:     os.Getenv("REDIS_ADDRESS"),
		RedisUsername: os.Getenv("REDIS_USERNAME"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       os.Getenv("REDIS_DB"),
		Port:          os.Getenv("PORT"),
	}
}
