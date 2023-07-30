FROM golang:1.20 AS build-stage-01

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o eros-api cmd/server/main.go

FROM alpine:latest AS production

# setup user
RUN groupmod -g 1000 users && \
  useradd -u 911 -U -d /data abc && \
  usermod -G users abc

WORKDIR /opt/app

COPY --from=build-stage-01 /app/eros-api .

EXPOSE 4000

# copy entrypoint
COPY entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]