package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() Config {
	godotenv.Load()

	return Config{
		Port: os.Getenv("PORT"),
	}
}