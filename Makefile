default: all
all: build docker-build docker-push

build:
	CGO_ENABLED=0 GOOS=linux go build

docker-build: build
	docker build -t jspc/weather-api .

docker-push: docker-build
	docker push jspc/weather-api
