package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token  string
	APIUrl string
	Debug  bool
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err.Error())
	}

	isProd := os.Getenv("DEBUG") == "false"

	return &Config{
		Token:  os.Getenv("TGBOTTOKEN"),
		APIUrl: os.Getenv("APIURL"),
		Debug:  !isProd,
	}
}
