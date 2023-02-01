package config

import (
	"fmt"
	"hexa-example-go/internal/pkg/logger"
	"hexa-example-go/internal/pkg/mongo"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type HexaExampleConfig struct {
	Logger logger.LoggerConfig
	Mongo  mongo.MongoConfig
}

/*
Load function first, tries to load config from .env file. If file does not exist .env is skipped.
Then parses envs from os.
*/
func Load() (config HexaExampleConfig, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Println("Could not load .env file. Skipping.")
	}

	err = env.Parse(&config)
	return
}
