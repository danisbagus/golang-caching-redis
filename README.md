# golang-caching-redis
Simple go app implements caching using redis

## Requirements

- [Golang](https://golang.org/) as main programming language.
- [Go Module](https://go.dev/blog/using-go-modules) for package management.
- [Docker-compose](https://docs.docker.com/compose/) for running Mysql and Redis Container.

## Setup

Create RabbitMQ container

```bash
docker-compose up
```

## Run the service

Get Go packages

```bash
go get .
```

Run application

```bash
go run cmd/main.go
`````
