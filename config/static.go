package config

import (
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

var Static = struct {
	HTTPServerPort string `env:"HTTP_SERVER_PORT" envDefault:"8080"`
	PGHost         string `env:"DB_HOST" envDefault:"dpg-cpg1rav79t8c73ec0ol0-a"`
	PGPort         string `env:"DB_PORT" envDefault:"5432"`
	PGUser         string `env:"DB_USER" envDefault:"weatheruser"`
	PGPassword     string `env:"DB_PASSWORD" envDefault:"JYXyRZqVU5V6u7JuvrcL7QJFz6h2DmK8"`
	DbName         string `env:"DB_NAME" envDefault:"weather_ukcq"`
}{}

func init() {
	err := env.Parse(&Static)
	if err != nil {
		logrus.WithError(err).Fatal("Could not load static configuration")
	}
}
