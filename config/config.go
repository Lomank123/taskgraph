package config

import "taskgraph/utils"

type Config struct {
	App      AppConfig
	Postgres PostgresConfig
}

type AppConfig struct {
	Port string
}

var Cfg Config

func LoadConfig() *Config {
	Cfg = Config{
		App: AppConfig{
			Port: utils.GetEnv("APP_PORT", "8000"),
		},
		Postgres: LoadPostgresConfig(),
	}
	return &Cfg
}
