package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds environment variables
type Config struct {
	RedisHost        string
	RedisPort        string
	RedisPassword    string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresSchema   string
	LogLevel         string
	CacheTTL         string
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	cfg := &Config{
		RedisHost:        os.Getenv("REDIS_HOST"),
		RedisPort:        os.Getenv("REDIS_PORT"),
		RedisPassword:    os.Getenv("REDIS_PASSWORD"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresSchema:   os.Getenv("POSTGRES_SCHEMA"),
		LogLevel:         os.Getenv("LOG_LEVEL"),
		CacheTTL:         os.Getenv("CACHE_TTL"),
	}
	return cfg
}
