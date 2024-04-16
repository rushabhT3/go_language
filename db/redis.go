// db/redis.go
package db

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func ConnectRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,  	// DB: 0, this refers to the Redis database number.
	})

	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
