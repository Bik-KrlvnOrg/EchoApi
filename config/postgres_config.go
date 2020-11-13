package config

import "os"

type PostgresConfig struct {
	Driver   string
	Host     string
	Name     string
	User     string
	Password string
	Port     string
}

func NewPostgresConfig() *PostgresConfig {
	return &PostgresConfig{}
}

func (pc *PostgresConfig) GetConfig() *PostgresConfig {
	return &PostgresConfig{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
	}
}
