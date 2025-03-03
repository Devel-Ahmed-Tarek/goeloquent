package goeloquent

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// ConnectRedis يتصل بخادم Redis
func ConnectRedis(addr string) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("❌ Failed to connect to Redis:", err)
	}
	log.Println("✅ Redis connected successfully!")
}

// CacheGet يسترجع قيمة من Redis
func CacheGet(key string) (string, error) {
	return RedisClient.Get(context.Background(), key).Result()
}

// CacheSet يضع قيمة في Redis لمدة معينة
func CacheSet(key string, value string, duration time.Duration) error {
	return RedisClient.Set(context.Background(), key, value, duration).Err()
}
