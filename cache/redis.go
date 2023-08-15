package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var (
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	ctx = context.Background()
)
