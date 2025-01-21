package config

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func (e *DatabaseConfig) DSN() string {
	dsn := fmt.Sprintf("postgres://%s", e.Username)

	if e.Password != "" {
		dsn = fmt.Sprintf("%s:%s", dsn, e.Password)
	}

	return fmt.Sprintf("%s@%s:%s/%s", dsn, e.Host, e.Port, e.Database)
}

type Config struct {
	DBSource      DatabaseConfig
	DBDestination DatabaseConfig
}

func NewConfig() Config {

	dbSourceConfig := DatabaseConfig{
		Host:     os.Getenv("DB_SOURCE_HOST"),
		Port:     os.Getenv("DB_SOURCE_PORT"),
		Username: os.Getenv("DB_SOURCE_USER"),
		Password: os.Getenv("DB_SOURCE_PASSWORD"),
		Database: os.Getenv("DB_SOURCE_NAME"),
	}

	dbDestinationConfig := DatabaseConfig{
		Host:     os.Getenv("DB_DESTINATION_HOST"),
		Port:     os.Getenv("DB_DESTINATION_PORT"),
		Username: os.Getenv("DB_DESTINATION_USER"),
		Password: os.Getenv("DB_DESTINATION_PASSWORD"),
		Database: os.Getenv("DB_DESTINATION_NAME"),
	}

	return Config{
		DBSource:      dbSourceConfig,
		DBDestination: dbDestinationConfig,
	}

}
