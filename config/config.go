package config

import (
	"testovoe/pkg/pg"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Server struct {
		Port int    `env:"SERVER_PORT"`
		Key  string `env:"KEY"`
	}
	MailerAddress string `env:"MAILER_ADDRESS"`
	Db            pg.PgCredentials
}

func InitConfig() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
