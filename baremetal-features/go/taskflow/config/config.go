package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)
//& types
	type Config struct {
		AppEnv string
		Port int
		DB DBConfig
	}

	type DBConfig struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
		SSLMode  string
	}
//& loads configuration from .env file and environment variables
	func Load() (*Config, error) {
		if err := godotenv.Load(); err != nil { //! inject my .env vars into the operating system's environment
			log.Println("No .env file found, using environment variables")
		}
		port, err := strconv.Atoi(getEnv("PORT", "8080"))
		if err != nil {
			return nil, fmt.Errorf("invalid PORT: %w", err)
		}

		dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
		if err != nil {
			return nil, fmt.Errorf("invalid DB_PORT: %w", err)
		}

		cfg := &Config{
			AppEnv: getEnv("APP_ENV", "development"),
			Port:   port,
			DB: DBConfig{
				Host:     getEnv("DB_HOST", "localhost"),
				Port:     dbPort,
				User:     getEnv("DB_USER", "postgres"),
				Password: getEnv("DB_PASSWORD", ""),
				Name:     getEnv("DB_NAME", "taskflow"),
				SSLMode:  getEnv("DB_SSLMODE", "disable"),
			},
		}

		return cfg, nil
	}
//& Data Source Name for the db connection
	func (db *DBConfig) DSN() string  {
		return fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			db.Host, db.Port, db.User, db.Password, db.Name, db.SSLMode,
		)
	}
//& Helper
	func getEnv(key, defaultValue string) string {
		if value := os.Getenv(key) ; value != "" {
			return value
		}
		return defaultValue
	}