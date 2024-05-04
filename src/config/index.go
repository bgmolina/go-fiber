package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	API_PORT    string
	PGSQL_URI   string
	API_VERSION string
	API_HOST    string
	ENV         string
}

func EnvFn() Env {
	// load .env file
	godotenv.Load(".env")

	// return os.Getenv(key)
	result := Env{
		API_PORT:    os.Getenv("API_PORT"),
		API_VERSION: os.Getenv("API_VERSION"),
		API_HOST:    os.Getenv("API_HOST"),
		ENV:         os.Getenv("ENV"),
	}

	return result
}
