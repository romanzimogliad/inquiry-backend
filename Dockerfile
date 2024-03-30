FROM alpine:latest

ADD ./bin/app /app
ADD ./config.yml /config.yml
CMD ["/app"]
