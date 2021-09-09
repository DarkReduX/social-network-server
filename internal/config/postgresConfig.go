package config

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type PostgresConfig struct {
	Port     string `env:"POSTGRES_PORT"`
	Host     string `env:"POSTGRES_HOST"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	DbName   string `env:"POSTGRES_DBNAME"`
	SslMode  string `env:"POSTGRES_SSL_MODE"`
}

func NewPostgresConfig() *PostgresConfig {
	cfg := PostgresConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.WithFields(log.Fields{
			"file": "postgresConfig.go",
			"func": "NewPostgresConfig()",
		}).Errorf("Unable to parse environment variables: %v", err)
	}
	return &cfg
}
