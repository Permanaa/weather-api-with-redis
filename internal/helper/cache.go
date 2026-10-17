package helper

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	redis *redis.Client
}

func NewCache(client *redis.Client) *Cache {
	return &Cache{
		redis: client,
	}
}

func (c *Cache) GetJson(ctx context.Context, key string, dest any) (bool, error) {
	value, err := c.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	if err := json.Unmarshal([]byte(value), dest); err != nil {
		return false, nil
	}

	return true, nil
}

func (c *Cache) SetJson(ctx context.Context, key string, value any, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.redis.Set(ctx, key, data, ttl).Err()
}
