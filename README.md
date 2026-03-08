# Go Service Template

A production-ready Go backend service template built with Gin, PostgreSQL, and JWT authentication.

## Features

- **Gin** HTTP framework with structured routing
- **PostgreSQL** with sqlc for type-safe database queries
- **JWT Authentication** middleware
- **Structured Logging** with zerolog
- **OpenAPI/Swagger** documentation
- **Docker** multi-stage build + docker-compose
- **GitHub Actions** CI pipeline (lint, test, build)
- Graceful shutdown
- Health check endpoint

## Quick Start

### Prerequisites

- Go 1.23+
- PostgreSQL 16+
- [sqlc](https://sqlc.dev/) (for regenerating DB code)
- [swag](https://github.com/swaggo/swag) (for regenerating Swagger docs)

### Run locally

```bash
# Copy environment config
cp .env.example .env

# Start PostgreSQL (via Docker)
docker compose up db -d

# Run migrations
make migrate-up

# Start the server
make run
```

### Run with Docker

```bash
docker compose up --build
```

### API Endpoints

| Method | Path                    | Auth | Description       |
|--------|-------------------------|------|-------------------|
| GET    | `/healthz`              | No   | Health check      |
| POST   | `/api/v1/auth/register` | No   | Register user     |
| POST   | `/api/v1/auth/login`    | No   | Login             |
| GET    | `/api/v1/users/me`      | Yes  | Get current user  |
| GET    | `/api/v1/users/:id`     | Yes  | Get user by ID    |
| GET    | `/swagger/*`            | No   | Swagger UI        |

### Development

```bash
make test          # Run tests
make lint          # Run linter
make sqlc          # Regenerate sqlc code
make swagger       # Regenerate Swagger docs
make build         # Build binary
```

## Project Structure

```
├── cmd/server/          # Application entrypoint
├── internal/
│   ├── config/          # Configuration loading
│   ├── handler/         # HTTP handlers
│   ├── middleware/       # Auth & logging middleware
│   ├── model/           # Request/response models
│   ├── repository/      # Database layer (sqlc generated)
│   └── service/         # Business logic
├── db/
│   ├── migrations/      # SQL migrations
│   ├── queries/         # sqlc query definitions
│   └── sqlc.yaml        # sqlc config
├── docs/                # Swagger docs (generated)
├── .github/workflows/   # CI pipeline
├── Dockerfile
├── docker-compose.yml
└── Makefile
```
