Weather API
==

Read the current, local weather from redis and return it to clients, over HTTP. This data is put into place by github.com/culture-trip/weather-job

## Building

This container can, ostensibly, be built with the usual and expected golang compilation toolchain, as per:

```bash
$ go get
$ go build
```

This will output the file `weather-api`. This file wont, though, be ready to drop into a container (unless you're on a linux box with cgo disabled); it'll just allow the tool to be run locally.

We provide a `Makefile` instead:

```bash
$ make clean weather-api
rm weather-api
CGO_ENABLED=0 GOOS=linux go build
```

This container can be pushed into the docker library as per:

```bash
$ make docker-build docker-push
```

Or:

```bash
$ make
```

Will push the container `example/weather-api`

## Deployment

This project should be deployed into kubernetes.


## Usage

This app is designed to be run as a docker container, as per:

`$ docker run -e REDIS_URL=localhost:6379 culture-trip/weather-api`

If connecting to a master/worker redis configuration over links, or over the same network, you can drop the `REDIS_URL` environment variable where the master is exposed via `REDIS_MASTER_SERVICE_URL`

## Gotchas

### `panic: dial tcp [::1]:6379: connect: connection refused`

This, or similar, means the tool can't access the specified redis
