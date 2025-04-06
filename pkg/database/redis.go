package database

import (
    "context"
    "fmt"
    "github.com/redis/go-redis/v9"
    "github.com/vanthang24803/mini/internal/config"
)

var RedisClient *redis.Client

func InitRedis(cfg *config.Config) error {
    client := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    // Test the connection
    ctx := context.Background()
    if err := client.Ping(ctx).Err(); err != nil {
        return fmt.Errorf("error connecting to redis: %v", err)
    }

    RedisClient = client
    return nil
}

func GetRedis() *redis.Client {
    return RedisClient
}