package database

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vanthang24803/mini/internal/config"
	"github.com/vanthang24803/mini/pkg/logger"
	"go.uber.org/zap"
)

var RedisClient *redis.Client

func InitRedis(cfg *config.Config) error {
	log := logger.GetLogger()
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
		DialTimeout:  time.Second * 2,
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 2,
		PoolSize:     10,
		PoolTimeout:  time.Second * 3,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Test connection
	if err := client.Ping(ctx).Err(); err != nil {
		log.Error("Redis connection failed", zap.Error(err))
		return fmt.Errorf("redis connection failed: %v", err)
	}

	log.Info("Redis connected successfully",
		zap.String("host", cfg.Redis.Host),
		zap.String("port", cfg.Redis.Port))

	RedisClient = client
	return nil
}

func GetRedis() *redis.Client {
	return RedisClient
}

// Helper functions for common Redis operations
func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

func Get(ctx context.Context, key string) (string, error) {
	return RedisClient.Get(ctx, key).Result()
}

func Del(ctx context.Context, keys ...string) error {
	return RedisClient.Del(ctx, keys...).Err()
}
