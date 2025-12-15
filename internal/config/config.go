package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DatabaseURL string
	JWTsecret string
}

func LoadConfig() Config {
	godotenv.Load()

	return Config{
		Port: os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTsecret: os.Getenv("JWT"),
	}
}