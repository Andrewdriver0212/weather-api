FROM alpine
MAINTAINER jspc

ADD weather-api /weather-api

ENTRYPOINT ["/weather-api"]
