package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

var (
	clock    Clock
	redisURL string
)

func main() {
	clock = realClock{}

	redisURL, err := envExpect("REDIS_URL")
	if err != nil {
		r, err2 := envExpect("REDIS_MASTER_SERVICE_HOST")
		if err2 != nil {
			panic(err)
		}

		redisURL = fmt.Sprintf("%s:6379", r)
	}

	log.Printf("Connecting to %s", redisURL)

	api := API{
		redis: redis.NewClient(&redis.Options{
			Addr: redisURL,
			DB:   0,
		}),
	}

	err = api.redis.Ping().Err()
	if err != nil {
		panic(err)
	}

	log.Print("Redis is connected")

	http.Handle("/", api)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func envExpect(key string) (val string, err error) {
	val = os.Getenv(key)
	if val == "" {
		err = fmt.Errorf("Value %q is not set in the environment", key)
	}

	return
}
