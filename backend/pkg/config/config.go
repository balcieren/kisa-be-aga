package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	UrlApiPort, UrlServicePort, UrlServiceDatabaseUrl string
}

func New() *Config {
	godotenv.Load("../../.env")
	return &Config{
		UrlApiPort:            os.Getenv("URL_API_PORT"),
		UrlServicePort:        os.Getenv("URL_SERVICE_PORT"),
		UrlServiceDatabaseUrl: os.Getenv("URL_SERVICE_DATABASE_URL"),
	}
}
