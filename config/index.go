package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type sEnv struct {
	API_PORT    string
	PGSQL_URI   string
	API_VERSION string
}

func EnvFn() sEnv {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	// return os.Getenv(key)
	result := sEnv{
		API_PORT:    os.Getenv("API_PORT"),
		PGSQL_URI:   os.Getenv("PGSQL_URI"),
		API_VERSION: os.Getenv("API_VERSION"),
	}

	return result
}
