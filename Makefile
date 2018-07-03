TAG ?= latest

default: all
all: weather-api docker-build docker-push

weather-api:
	CGO_ENABLED=0 GOOS=linux go build

docker-build: weather-api
	docker build -t jspc/weather-api:latest -t jspc/weather-api:$(TAG) .

docker-push: docker-build
	docker push jspc/weather-api:latest
	docker push jspc/weather-api:$(TAG)

clean:
	-rm weather-api

.PHONY: deploy
deploy:
	kubectl apply -f deploy/k8s/redis.yaml
	kubectl apply -f deploy/k8s/weather-api.yaml

	kubectl set image deployment/weather-api weather-api=jspc/weather-api:$(TAG)
