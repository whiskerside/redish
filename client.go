package redish

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func client() *redis.Client {
	db, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))
	address := fmt.Sprintf("%s:%s", getEnv("REDIS_HOST", "127.0.0.1"), getEnv("REDIS_PORT", "6379"))
	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       db,
	})
}
