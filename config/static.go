package config

import (
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

var Static = struct {
	HTTPServerPort string `env:"HTTP_SERVER_PORT" envDefault:":8080"`
	PGHost         string `env:"DB_HOST" envDefault:"postgres"`
	PGPort         string `env:"DB_PORT" envDefault:":5432"`
	PGUser         string `env:"DB_USER" envDefault:"postgres"`
	PGPassword     string `env:"DB_PASSWORD" envDefault:"Aa123456!"`
	DbName         string `env:"DB_NAME" envDefault:"weather"`
}{}

func init() {
	err := env.Parse(&Static)
	if err != nil {
		logrus.WithError(err).Fatal("Could not load static configuration")
	}
}
