package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"nats-streaming-web/internal/logger"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(redisURL string) *RedisClient {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		logger.ErrorLogger.Fatalf("Ошибка при парсинге Redis URL: %v", err)
	}

	client := redis.NewClient(opts)
	return &RedisClient{client: client}
}

func (rc *RedisClient) Set(ctx context.Context, key string, value interface{}) error {
	err := rc.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		logger.ErrorLogger.Printf("Ошибка при записи в Redis: %v", err)
	}
	return err
}

func (rc *RedisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := rc.client.Get(ctx, key).Result()
	if err != nil {
		logger.ErrorLogger.Printf("Ошибка при чтении из Redis: %v", err)
		return "", err
	}
	return val, nil
}
