package repository

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

// InitRedisClient initializes the Redis client and establishes a connection.
func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // Password (if any)
		DB:       0,                // Redis database number
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Redis connection failed: %v\n", err)
	}
	log.Printf("Redis connected: %s\n", pong)
}

// GetRedisClient returns the Redis client instance.
func GetRedisClient() *redis.Client {
	return client
}
