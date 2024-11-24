package redisdb

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	User     string
	Password string
	Host     string
	Name     string
}

func OpenRDB(conf Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	// Ping the Redis server
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		// log.Fatalf("Could not connect to Redis: %v", err)
		return nil, err
	}
	fmt.Println("connected to Redis !!", pong)
	return rdb, nil
}
