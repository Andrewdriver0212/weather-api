package main

import (
	"github.com/go-redis/redis"
)

//  Get(key string) *redis.StringCmd
//	Ping() *redis.StatusCmd

type dummyRedis struct {
	body string
	err  error
}

func (r dummyRedis) Get(_ string) *redis.StringCmd {
	return redis.NewStringResult(r.body, r.err)
}

func (r dummyRedis) Ping() *redis.StatusCmd {
	return redis.NewStatusResult("1", nil)
}
