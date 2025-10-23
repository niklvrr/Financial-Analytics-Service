package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var (
	dbUserEmptyError = errors.New("DB User is Empty")
	dbNameEmptyError = errors.New("DB Name is Empty")
	envLoadError     = errors.New(".env load Error")
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Env string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	URL      string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, envLoadError
	}

	cfg := &Config{
		App: AppConfig{
			Env: getEnv("APP_ENV", "local"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", ""),
			URL:      getEnv("DB_URL", ""),
		},
	}

	if cfg.Database.URL == "" {
		if cfg.Database.User == "" {
			return nil, dbUserEmptyError
		}
		if cfg.Database.Name == "" {
			return nil, dbNameEmptyError
		}
		cfg.Database.URL = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.Name,
		)
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
