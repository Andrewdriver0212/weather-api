FROM alpine
MAINTAINER jspc

EXPOSE 8000

ADD weather-api /weather-api

ENTRYPOINT ["/weather-api"]
