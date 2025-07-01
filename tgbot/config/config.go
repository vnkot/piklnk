package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %v", err.Error()))
	}

	return &Config{
		Token: os.Getenv("TGBOTTOKEN"),
	}
}
