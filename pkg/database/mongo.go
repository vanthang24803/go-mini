package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/vanthang24803/mini/internal/config"
	"github.com/vanthang24803/mini/pkg/constant"
	"github.com/vanthang24803/mini/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoDB *mongo.Database
	Client  *mongo.Client
)

func InitMongoDB(cfg *config.Config) error {
	log := logger.GetLogger()
	maxRetries := 5
	var lastErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		collections := []string{
			constant.COLLECTION_USER,
			constant.COLLECTION_TOKEN,
		}

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDB.URI))
		if err != nil {
			lastErr = err
			log.Error(fmt.Sprintf("mongoDB connection attempt %d failed: %v", attempt, err))
			if attempt == maxRetries {
				return fmt.Errorf("failed to connect after %d attempts: %v", maxRetries, lastErr)
			}
			time.Sleep(2 * time.Second)
			continue
		}

		if err := client.Ping(ctx, nil); err != nil {
			lastErr = err
			log.Error(fmt.Sprintf("MongoDB ping attempt %d failed: %v", attempt, err))
			if attempt == maxRetries {
				return fmt.Errorf("failed to ping after %d attempts: %v", maxRetries, lastErr)
			}
			time.Sleep(2 * time.Second)
			continue
		}

		Client = client
		MongoDB = client.Database(cfg.MongoDB.Database)

		for _, name := range collections {
			err := MongoDB.CreateCollection(ctx, name)
			if err != nil {
				if strings.Contains(err.Error(), "already exists") {
					continue
				}
				lastErr = err
				log.Error(fmt.Sprintf("failed to create collection %s: %v", name, err))
				if attempt == maxRetries {
					return fmt.Errorf("failed to create collections after %d attempts: %v", maxRetries, lastErr)
				}
				time.Sleep(2 * time.Second)
				continue
			}
		}

		log.Info(fmt.Sprintf("Connected to MongoDB ✔️ at %d", attempt))
		return nil
	}

	return lastErr
}

func CloseMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return Client.Disconnect(ctx)
}

func GetCollection(name string) *mongo.Collection {
	return MongoDB.Collection(name)
}
