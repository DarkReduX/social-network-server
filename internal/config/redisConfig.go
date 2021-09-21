package config

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type RedisConfig struct {
	Password string `env:"REDIS_PASSWORD"`
	Addr     string `env:"REDIS_URI"`
	Username string `env:"REDIS_USERNAME"`
	DB       int    `env:"REDIS_DB"`
}

func NewRedisConfig() *RedisConfig {
	cfg := RedisConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.WithFields(log.Fields{
			"file": "redisConfig.go",
			"func": "NewRedisConfig()",
		}).Errorf("Unable to parse environment variables: %v", err)
	}
	return &cfg
}
