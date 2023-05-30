package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	keyPrefix     = "shorturl:"
	expirationTTL = time.Hour * 24
)

type RedisRepository struct {
	client *redis.Client
}

type RedisRepositoryImpl interface {
	SaveShortURL(ctx context.Context, shortURLID string, longURL string) error
	GetLongURL(ctx context.Context, shortURLID string) (string, error)
}

func NewRedisRepository() RedisRepositoryImpl {
	return &RedisRepository{
		client: GetRedisClient(),
	}
}

// SaveShortURL saves the short URL and its associated long URL in Redis.
func (r *RedisRepository) SaveShortURL(ctx context.Context, shortURLID string, longURL string) error {
	key := fmt.Sprintf("%s%s", keyPrefix, shortURLID)
	value, err := json.Marshal(longURL)
	if err != nil {
		return err
	}

	err = r.client.Set(ctx, key, value, expirationTTL).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetLongURL retrieves the long URL associated with the given short URL from Redis.
func (r *RedisRepository) GetLongURL(ctx context.Context, shortURLID string) (string, error) {
	key := fmt.Sprintf("%s%s", keyPrefix, shortURLID)

	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("Short URL not found")
		}
		return "", err
	}

	var longURL string
	err = json.Unmarshal([]byte(result), &longURL)
	if err != nil {
		return "", err
	}

	return longURL, nil
}
