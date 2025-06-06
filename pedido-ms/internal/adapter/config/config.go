package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	MongoConnectionUrl string
	ServerPort         string
	RabbitHost         string
}

func New() (*AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	appConfig := AppConfig{
		MongoConnectionUrl: os.Getenv("MONGO_CONNECTION_URL"),
		ServerPort:         os.Getenv("SERVER_PORT"),
		RabbitHost:         os.Getenv("RABBIT_HOST"),
	}

	return &appConfig, nil
}
