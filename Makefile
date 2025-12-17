.PHONY: up down build migrate run

up:
	docker compose up -d

down:
	docker compose down

build:
	go build -o bin/api .

migrate-up:
	migrate -path migrations -database "postgres://taskapi:taskapi@localhost:5432/task_api?sslmode=disable" up

run:
	go run .