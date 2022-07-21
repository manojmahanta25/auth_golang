package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

func GetHost() string {
	host := os.Getenv("APP_PROXY_HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	return host
}

func GetPort() string {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	return port
}
