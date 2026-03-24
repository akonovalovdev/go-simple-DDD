.PHONY: build run migrate-up migrate-down migrate-create lint test clean deps

BINARY_NAME=server
DATABASE_URL?=postgres://postgres:postgres@localhost:5432/<base-template>?sslmode=disable
MIGRATIONS_DIR=migrations

deps:
	go mod download
	go mod tidy

migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" down

migrate-create:
	@read -p "Enter migration name: " name; \
	goose -dir $(MIGRATIONS_DIR) create $$name sql

migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" status

run:
	go run cmd/server/main.go

build:
	go build -o bin/$(BINARY_NAME) cmd/server/main.go

lint:
	golangci-lint run ./...

test:
	go test -v -race ./...

test-coverage:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

docker-build:
	docker build -t <template> .

docker-run:
	docker-compose up -d

docker-down:
	docker-compose down
