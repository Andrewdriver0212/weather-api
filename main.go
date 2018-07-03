package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

var (
	clock Clock
)

func main() {
	redisURL, err := envExpect("REDIS_URL")
	if err != nil {
		panic(err)
	}

	log.Printf("Connecting to %s", redisURL)

	clock = realClock{}

	api := API{
		redis: redis.NewClient(&redis.Options{
			Addr: redisURL,
			DB:   0,
		}),
	}

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
