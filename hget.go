package redish

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func HGetString(key, field string) string {
	rdb := client()
	defer rdb.Close()

	v, err := rdb.HGet(context.Background(), key, field).Result()
	if err == redis.Nil {
		log.Printf("the key [%s] doesn't exist", key)
		return ""
	}
	if err != nil {
		log.Println("something happended,", err.Error())
	}
	return v
}

func HSetString(key string, vals ...string) bool {
	rdb := client()
	defer rdb.Close()
	_, err := rdb.HSet(context.Background(), key, vals).Result()
	return err == nil
}
