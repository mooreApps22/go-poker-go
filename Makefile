NAME := poker-server

.PHONY: all build run clean docker-up docker-down docker-clean

all: build

build:
	go build -o $(NAME) ./cmd/server

run:
	go run ./cmd/server

clean:
	rm -f $(NAME)

docker-up:
	docker compose up --build

docker-down:
	docker compose down

docker-clean:
	docker compose down -v

