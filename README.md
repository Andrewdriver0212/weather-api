Weather API
==

Read the current, local weather from redis and return it to clients, over HTTP. This data is put into place by github.com/jspc/weather-job, and lives on github.com/jspc/tf-ct-tech-test.

## Building

This repo comes with a Makefile for convenience, but there's nothing complex about the build- a `go build` and `docker build .` are all that's needed.

## Usage

This app is designed to be run as a docker container, as per:

$ docker run -e REDIS_URL=localhost:6379 jspc/weather-api

## Gotchas

### `panic: dial tcp [::1]:6379: connect: connection refused`

This, or similar, means the tool can't access the specified redis
