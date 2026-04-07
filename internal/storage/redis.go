package storage

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func InitRedis(host string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: host,
		Password: "",
		DB: 0,
		Protocol: 2,
	})
	if err := client.Ping(context.Background()).Err() ;err != nil {
		panic(fmt.Errorf("%v", err))
	}

	return client
}
