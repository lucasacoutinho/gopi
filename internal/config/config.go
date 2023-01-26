package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr        string
	Debug       bool
	DatabaseURL string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	debug, _ := strconv.ParseBool(os.Getenv("APP_DEBUG"))

	return &Config{
		Addr:        fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")),
		Debug:       debug,
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
