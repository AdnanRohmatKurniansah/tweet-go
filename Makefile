include .env
export

DB_URL=mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?charset=utf8mb4&parseTime=True&loc=Asia%2FJakarta

.PHONY: run build docker-up docker-down \
        migrate-up migrate-down migrate-down-all \
        migrate-create migrate-version migrate-force

# App
run:
	go run ./cmd/main.go

dev:
	air

build:
	go build -o bin/app ./cmd/main.go

# Docker
docker-up:
	docker compose up -d

docker-down:
	docker compose down

# Migration
migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down 1

migrate-down-all:
	migrate -path migrations -database "$(DB_URL)" down

migrate-version:
	migrate -path migrations -database "$(DB_URL)" version

migrate-create:
	@read -p "Migration Name: " name; \
	migrate create -ext sql -dir migrations -seq $$name

migrate-force:
	@read -p "Force to version: " v; \
	migrate -path migrations -database "$(DB_URL)" force $$v

migrate-fresh:
	migrate -path migrations -database "$(DB_URL)" down -all
	migrate -path migrations -database "$(DB_URL)" up