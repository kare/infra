package infra

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// RedisClient is a redis.Client.
type RedisClient struct {
	*redis.Client
}

// NewRedisClient creates a new RedisClient.
func NewRedisClient(redisURL string) (*RedisClient, error) {
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("infra: error while parsing Redis URL: %w", err)
	}
	return &RedisClient{
		Client: redis.NewClient(options),
	}, nil
}

// Shutdown releases used resouces.
func (r *RedisClient) Shutdown(ctx context.Context) error {
	status := r.ShutdownSave(ctx)
	if err := status.Err(); err != nil {
		return fmt.Errorf("infra: error while disconnecting from Redis: %w", err)
	}
	return nil
}
