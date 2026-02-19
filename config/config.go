package config

import "taskgraph/utils"

type Config struct {
	App AppConfig
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
	}
	return &Cfg
}