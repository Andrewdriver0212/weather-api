package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestAPI_ServeHTTP(t *testing.T) {
	clock = frozenClock{}

	for _, test := range []struct {
		name         string
		redis        redisClient
		route        string
		method       string
		expectBody   string
		expectStatus int
	}{
		{"Options request", dummyRedis{}, "/", "OPTIONS", "", 200},
		{"POST request", dummyRedis{}, "/", "POST", `{"message":"Method: \"POST\" not allowed","status":405}`, 405},
		{"DELETE request", dummyRedis{}, "/", "DELETE", `{"message":"Method: \"DELETE\" not allowed","status":405}`, 405},
		{"Healthcheck endpoint", dummyRedis{}, "/healthcheck", "GET", "ok", 200},
		{"Broken redis connection", dummyRedis{err: fmt.Errorf("some error")}, "/", "GET", `{"message":"No weather data found","status":500}`, 500},
		{"No data in redis", dummyRedis{}, "/", "GET", `{"message":"No weather data found","status":500}`, 500},
		{"Happy path", dummyRedis{body: "{}"}, "/", "GET", "{}", 200},
		{"No such path", dummyRedis{}, "/nonsuch", "GET", `{"message":"\"/nonsuch\" not found","status":404}`, 404},
	} {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(test.method, test.route, nil)

			a := API{
				redis: test.redis,
			}

			a.ServeHTTP(w, r)
			resp := w.Result()

			body, _ := ioutil.ReadAll(resp.Body)
			if test.expectBody != string(body) {
				t.Errorf("expected %q, received %q", test.expectBody, string(body))
			}

			if test.expectStatus != resp.StatusCode {
				t.Errorf("expected %d, received %d", test.expectStatus, resp.StatusCode)
			}
		})
	}
}
