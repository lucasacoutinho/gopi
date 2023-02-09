package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	AppName        string
	Root           string
	Addr           string
	Debug          bool
	DatabaseDriver string
	DatabaseURL    string
	Logger         *zap.SugaredLogger
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	debug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	appName := os.Getenv("APP_NAME")
	log, err := initLog(appName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &Config{
		AppName:        appName,
		Root:           dir,
		Addr:           fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")),
		Debug:          debug,
		DatabaseDriver: os.Getenv("DATABASE_DRIVER"),
		DatabaseURL:    os.Getenv("DATABASE_URL"),
		Logger:         log,
	}
}

func initLog(service string) (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.InitialFields = map[string]interface{}{
		"service": service,
	}

	log, err := config.Build()
	if err != nil {
		return nil, err
	}

	return log.Sugar(), nil
}
