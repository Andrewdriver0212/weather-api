package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type API struct {
    redis redisClient
}

func (a API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("content-type", "application/json")

    w.Header().Set("access-control-allow-origin", "*")
    w.Header().Set("access-control-allow-headers", "origin, content-type, accept")
    w.Header().Set("access-control-allow-methods", "POST")

    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)

        return
    }

    if r.Method != http.MethodGet {
        Error(w, nil, http.StatusMethodNotAllowed, fmt.Sprintf("Method: %q not allowed", r.Method))

        return
    }

    switch r.URL.Path {
    case "/":
        r, err := a.redis.Get("weather").Bytes()
        if err != nil || len(r) == 0 {
            Error(w, err, http.StatusInternalServerError, fmt.Sprintf("No weather data found"))

            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write(r)

        return

    default:
        Error(w, nil, http.StatusNotFound, fmt.Sprintf("%q not found", r.URL.Path))
    }

}

func Error(w http.ResponseWriter, e error, status int, message string) {
    body := map[string]interface{}{
        "status":  status,
        "message": message,
    }

    output, err := json.Marshal(body)
    if err != nil {
        log.Printf("%+v", err)
        w.WriteHeader(http.StatusInternalServerError)

        return
    }

    w.WriteHeader(status)
    w.Write(output)

    body["error"] = e
    body["time"] = clock.Now()

    output, err = json.Marshal(body)
    if err != nil {
        log.Printf("%+v", err)

        return
    }

    fmt.Println(string(output))
}
