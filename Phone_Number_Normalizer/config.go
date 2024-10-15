package main

import "github.com/caarlos0/env/v11"

type Config struct {
	dbConnectionUrl string `env:"POSTGRES_CONN_URL"`
}

func ParseConfig() *Config {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
