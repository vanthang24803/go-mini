package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/vanthang24803/mini/pkg/util"
)

func New() *Config {
	godotenv.Load()

	return &Config{
		AppName: os.Getenv(util.APP_NAME),
		Server: ServerConfig{
			Port: ":" + getEnvOrDefault(util.APP_PORT, "3000"),
		},
		Logger: LoggerConfig{
			Level:      "info",
			OutputPath: "logs",
			MaxAge:     30,
			Production: os.Getenv(util.ENV) == util.PRODUCTION,
		},
		Database: DatabaseConfig{
			Host:     os.Getenv(util.DB_HOST),
			Port:     os.Getenv(util.DB_PORT),
			Username: os.Getenv(util.DB_USERNAME),
			Password: os.Getenv(util.DB_PASSWORD),
			Name:     os.Getenv(util.DB_NAME),
		},
		Redis: RedisConfig{
			Host: os.Getenv(util.REDIS_HOST),
			Port: os.Getenv(util.REDIS_PORT),
		},
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
