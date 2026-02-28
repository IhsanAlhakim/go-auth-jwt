package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName         string
	JWTSigKey       string
	TokenCookieName string

	DBName     string
	DBUsername string
	DBPassword string
	DBAddr     string

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
		AppName:         os.Getenv("APP_NAME"),
		JWTSigKey:       os.Getenv("JWT_SIGNATURE_KEY"),
		TokenCookieName: os.Getenv("TOKEN_COOKIE_NAME"),
		DBName:          os.Getenv("DB_NAME"),
		DBUsername:      os.Getenv("DB_USERNAME"),
		DBPassword:      os.Getenv("DB_PASSWORD"),
		DBAddr:          os.Getenv("DB_ADDRESS"),
		Port:            port,
	}
}
