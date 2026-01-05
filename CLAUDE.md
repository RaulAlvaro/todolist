# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A Todo List REST API built with Go, using Gin web framework, GORM ORM with PostgreSQL, and Google Wire for dependency injection. The application follows a clean architecture pattern with clear separation between handlers, services, repositories, and domain entities.

## Development Commands

### Running the Application
```bash
go run cmd/*.go
```

The server will start on the port specified in the `PORT` environment variable (defaults from `.env` file).

### Dependency Injection with Wire

This project uses Google Wire for compile-time dependency injection. After modifying dependency providers or the wire configuration:

```bash
# Generate wire_gen.go from wire.go
go generate ./cmd

# Or run wire directly
wire ./cmd
```

Wire configuration is in [cmd/wire.go](cmd/wire.go) and generates [cmd/wire_gen.go](cmd/wire_gen.go).

### Database Setup

The application uses GORM AutoMigrate, so database tables are created automatically on startup. Ensure PostgreSQL is running and configure the connection via `.env`:

```bash
# Copy example environment file
cp .env.example .env

# Edit .env with your database credentials
# DB_URL format: host=localhost user=postgres password=postgres dbname=todolist port=5432 sslmode=disable
```

Default connection (if `DB_URL` is not set): `host=localhost user=postgres password=postgres dbname=todolist port=5432 sslmode=disable`

### Managing Dependencies
```bash
# Add a new dependency
go get github.com/package/name

# Tidy dependencies
go mod tidy
```

## Architecture

### Layered Structure

The codebase follows a clean architecture with these layers:

1. **Handler Layer** ([internal/handler/](internal/handler/)) - HTTP handlers using Gin
   - Receives HTTP requests, validates input, calls services
   - Uses [pkg/response/response.go](pkg/response/response.go) for standardized JSON responses
   - All responses follow the format: `{success: bool, message: string, data: any, error: string}`

2. **Service Layer** ([internal/service/](internal/service/)) - Business logic
   - Contains domain validations and orchestration
   - Handles transactions when multiple operations must be atomic
   - Example: `CreateWithAudit` in [internal/service/todo_service.go](internal/service/todo_service.go) demonstrates transaction usage

3. **Repository Layer** ([internal/repository/](internal/repository/)) - Data access
   - Implements repository interfaces defined in the domain package
   - All methods accept `context.Context` for cancellation/timeout support
   - Returns interface types, not concrete implementations

4. **Domain Layer** ([internal/domain/](internal/domain/)) - Core entities and interfaces
   - Contains entity definitions (e.g., `Todo` struct with GORM tags)
   - Defines repository interfaces (e.g., `TodoRepository`)
   - No external dependencies except GORM for model definitions

### Dependency Injection Flow

Wire automatically generates the dependency graph in [cmd/wire_gen.go](cmd/wire_gen.go) based on providers in [cmd/wire.go](cmd/wire.go):

```
Config → Database → Repository → Service → Handler → App
```

The `initializeApp()` function in [cmd/main.go](cmd/main.go) uses Wire to construct all dependencies.

### Key Patterns

**Repository Pattern**: Repositories return interface types. This allows the service layer to depend on abstractions, not implementations:
```go
func NewTodoRepository(db *gorm.DB) domain.TodoRepository
```

**Transaction Handling**: When using transactions in services, create a new repository instance with the transaction connection:
```go
s.db.Transaction(func(tx *gorm.DB) error {
    txRepo := repository.NewTodoRepository(tx)
    // Use txRepo instead of s.todoRepo
})
```

**Context Propagation**: All repository and service methods accept `context.Context` from the HTTP request for proper cancellation and timeout handling.

## API Routes

All routes are prefixed with `/api/v1` and defined in [cmd/main.go](cmd/main.go):

- `GET /api/v1/todos` - Get all todos
- `GET /api/v1/todos/:id` - Get todo by ID
- `POST /api/v1/todos` - Create new todo (requires `content` field)

## Configuration

Configuration is loaded via environment variables in [config/config.go](config/config.go):

- `PORT` - Server port
- `DB_URL` - PostgreSQL connection string

Uses `joho/godotenv` to load from `.env` file automatically.
