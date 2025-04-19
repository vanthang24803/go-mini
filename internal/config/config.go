package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/vanthang24803/mini/pkg/constant"
)

func New() *Config {
	godotenv.Load()

	return &Config{
		AppName: os.Getenv(constant.APP_NAME),
		Server: ServerConfig{
			Port: ":" + getEnvOrDefault(constant.APP_PORT, "3000"),
		},
		Logger: LoggerConfig{
			Level:      "info",
			OutputPath: "logs",
			MaxAge:     30,
			Production: os.Getenv(constant.ENV) == "production",
		},
		MongoDB: MongoConfig{
			URI:      os.Getenv(constant.MONGODB_URI),
			Database: os.Getenv(constant.MONGODB_DATABASE),
		},
		Redis: RedisConfig{
			Host:     getEnvOrDefault(constant.REDIS_HOST, "localhost"),
			Port:     getEnvOrDefault(constant.REDIS_PORT, "6379"),
			Password: os.Getenv(constant.REDIS_PASSWORD),
			DB:       0,
		},
		Minio: MinioConfig{
			Endpoint:        os.Getenv(constant.MINIO_ENDPOINT),
			AccessKeyID:     os.Getenv(constant.MINIO_ACCESS_KEY_ID),
			SecretAccessKey: os.Getenv(constant.MINIO_SECRET_ACCESS_KEY),
			BucketName:      os.Getenv(constant.MINIO_BUCKET_NAME),
			UseSSL:          os.Getenv(constant.MINIO_USE_SSL) == "true",
		},
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
