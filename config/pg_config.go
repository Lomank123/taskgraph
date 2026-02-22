package config

import "taskgraph/utils"

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	DSN      string
}

func LoadPostgresConfig() PostgresConfig {
	host := utils.GetEnv("POSTGRES_HOST", "localhost")
	port := utils.GetEnv("POSTGRES_PORT", "5436")
	user := utils.GetEnv("POSTGRES_USER", "postgres")
	password := utils.GetEnv("POSTGRES_PASSWORD", "postgres")
	name := utils.GetEnv("POSTGRES_DB", "postgres")
	sslmode := utils.GetEnv("POSTGRES_SSLMODE", "disable")
	dsn := utils.BuildDBDSN(
		host,
		port,
		user,
		password,
		name,
		sslmode,
	)

	return PostgresConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     name,
		SSLMode:  sslmode,
		DSN:      dsn,
	}
}
