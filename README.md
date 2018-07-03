Weather API
==

Read the current, local weather from redis and return it to clients, over HTTP. This data is put into place by github.com/jspc/weather-job, and lives on github.com/jspc/tf-ct-tech-test.

## Building

This container can, ostensibly, be built with the usual and expected golang compilation toolchain, as per:

```bash
$ go get
$ go build
```

This will output the file `weather-api`. This file wont, though, be ready to drop into a container (unless you're on a linux box with cgo enabled); it'll just allow the tool to be run locally.

We provide a `Makeile` instead:

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

This will, probably, fail for anybody who doesn't have access to my docker account. Thus:

```bash
$ make USER=example
```

Will push the container `example/weather-api`

## Deployment

This project deploys into kubernetes.

### Pre-deploy configuration

For k8s clusters which are not accessible over localhost you will need:

 1. A kubernetes client admin.conf file
 1. An environment variable (`$KUBECONFIG`) pointing to this file.

For environments built with `tf-ct-tech-test` this exists in `.secrets/admin.conf`

### Deploying

Assuming the pre-reqs are configured:

```bash
make deploy
```

Will build a redis master/slave cluster and will configure `weather-api` to use it


## Usage

This app is designed to be run as a docker container, as per:

$ docker run -e REDIS_URL=localhost:6379 jspc/weather-api

## Gotchas

### `panic: dial tcp [::1]:6379: connect: connection refused`

This, or similar, means the tool can't access the specified redis
