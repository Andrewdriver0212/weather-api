package main

import (
	"github.com/go-redis/redis"
)

type redisClient interface {
	Get(key string) *redis.StringCmd
	Ping() *redis.StatusCmd
}
