package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port string
	DB   DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Config{
		Port: port,
		DB:   loadDBConfig(),
	}
}

func loadDBConfig() DBConfig {
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	if dbConfig.Host == "" {
		dbConfig.Host = "localhost"
	}
	if dbConfig.Port == "" {
		dbConfig.Port = "5432"
	}
	if dbConfig.SSLMode == "" {
		dbConfig.SSLMode = "disable"
	}

	return dbConfig
}

func (c DBConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Name, c.SSLMode,
	)
}
