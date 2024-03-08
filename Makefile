all: build start

start:
	./bin/init

dev:
	go run internal/main.go

build:
	go build -o bin/init internal/main.go