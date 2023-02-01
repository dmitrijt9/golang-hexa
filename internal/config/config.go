package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

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

type LoggerConfig struct {
	Level           string `env:"LOG_LEVEL" envDefault:"info"`
	ConsoleEnabled  bool   `env:"LOG_CONSOLE_IS_ENABLED" envDefault:"true"`
	FilebeatEnabled bool   `env:"LOG_FILEBEAT_IS_ENABLED" envDefault:"false"`
	FilebeatUrl     string `env:"LOG_FILEBEAT_URL"`
	FilebeatIndex   string `env:"LOG_FILEBEAT_INDEX"`
	FileBeatAppName string `env:"LOG_FILEBEAT_APPNAME"`
}

type ServerConfig struct {
	Host string `env:"HOST" envDefault:"localhost"`
	Port string `env:"PORT" envDefault:"8080"`
}

type MongoConfig struct {
	Host     string `env:"MONGO_HOST" envDefault:"localhost"`
	Port     string `env:"MONGO_PORT" envDefault:"27017"`
	User     string `env:"MONGO_USER" envDefault:"user"`
	Password string `env:"MONGO_PASSWORD" envDefault:"pass"`
	Database string `env:"MONGO_DATABASE" envDefault:"test"`
}

type HexaExampleConfig struct {
	Logger LoggerConfig
	Mongo  MongoConfig
	Server ServerConfig
}
