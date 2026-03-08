.PHONY: build run test lint docker-up docker-down migrate-up migrate-down sqlc swagger

# Build the application
build:
	go build -o bin/server ./cmd/server

# Run the application locally
run:
	go run ./cmd/server

# Run tests
test:
	go test -v -race ./...

# Run tests with coverage
test-cover:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run linter
lint:
	golangci-lint run ./...

# Start all services with Docker Compose
docker-up:
	docker compose up --build -d

# Stop all services
docker-down:
	docker compose down

# Run database migrations (requires golang-migrate)
migrate-up:
	migrate -path db/migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path db/migrations -database "$(DATABASE_URL)" down

# Generate sqlc code
sqlc:
	cd db && sqlc generate

# Generate Swagger docs (requires swag CLI)
swagger:
	swag init -g cmd/server/main.go -o docs
