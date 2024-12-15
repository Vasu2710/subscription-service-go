package database

import (
	"context"
	"log"
	"strings"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var Ctx = context.Background()

// InitRedis initializes the Redis client
func InitRedis(redisURL string) {
	// Remove "redis://" prefix if present
	redisAddr := strings.TrimPrefix(redisURL, "redis://")

	RedisClient = redis.NewClient(&redis.Options{
		Addr: redisAddr, // Pass cleaned-up address
	})

	// Test Redis connection
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully!")
}
