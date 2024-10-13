package Lib

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func TestRedis(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("Redis client connect failed: ", err)
	}
	fmt.Println("Redis client connect succeeded")
	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal("Redis client close failed: ", err)
		}
	}(client)

	k := "key"
	if err := client.SetNX(k, "value", time.Second*10).Err(); err != nil {
		log.Fatal("Redis client setNX failed: ", err)
	}

	value, err := client.Get(k).Result()
	if err != nil {
		log.Fatal("Redis client get failed: ", err)
	}
	fmt.Println("Redis client get succeeded: ", value)
}
