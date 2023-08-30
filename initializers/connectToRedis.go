package initializers

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var REDISCLIENT *redis.Client

func ConnectToRedis() {
	// Initialize a Redis client
	REDISCLIENT = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // Password (if needed)
		DB:       0,                // Default DB
	})

	// Test the connection
	_, err := REDISCLIENT.Ping(context.Background()).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
}
