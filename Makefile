SERVER-NAME := poker-server
CLIENT-NAME := poker-client

.PHONY: all build run clean docker-up docker-down docker-clean

all: build

build:
	go build -o $(SERVER-NAME) ./cmd/server
	go build -o $(CLIENT-NAME) ./cmd/client

run:
	go run ./cmd/server

clean:
	rm -f $(SERVER-NAME)
	rm -f $(CLIENT-NAME)

docker-up:
	docker compose up --build

docker-down:
	docker compose down

docker-clean:
	docker compose down -v

test-poker:
	go test -v ./internal/poker
