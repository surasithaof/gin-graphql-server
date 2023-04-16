package app

import (
	"log"
	"surasithit/gin-graphql-server/adapters/db"
	"surasithit/gin-graphql-server/adapters/httpserver"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HttpServer httpserver.Config
	Database   db.Config
}

func LoadConfig() (*Config, error) {
	AppConfig := &Config{}
	err := envconfig.Process("", AppConfig)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return AppConfig, nil
}
