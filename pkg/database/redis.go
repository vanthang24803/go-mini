package database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vanthang24803/mini/internal/config"
	"github.com/vanthang24803/mini/pkg/logger"
	"go.uber.org/zap"
)

var redisClient *redis.Client

type Redis interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string, dest any) error
	Del(ctx context.Context, keys ...string) error
}

type RedisService struct {
	client *redis.Client
}

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Error("Redis connection failed", zap.Error(err))
		return fmt.Errorf("redis connection failed: %w", err)
	}

	log.Info("Redis connected successfully!")
	redisClient = client
	return nil
}

func NewRedisService() Redis {
	return &RedisService{client: redisClient}
}

func (r *RedisService) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, bytes, expiration).Err()
}

func (r *RedisService) Get(ctx context.Context, key string, dest any) error {
	data, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), dest)
}

func (r *RedisService) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}
