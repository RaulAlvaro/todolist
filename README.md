# Todo List API

A REST API for managing todo items, built with Go using clean architecture principles.

## Features

- RESTful API endpoints for todo management
- Clean architecture with dependency injection
- PostgreSQL database with GORM ORM
- Automatic database migrations
- Standardized JSON responses
- Context-aware operations with timeout/cancellation support

## Tech Stack

- **Framework**: [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- **ORM**: [GORM](https://gorm.io/) - Object-relational mapping
- **Database**: PostgreSQL
- **Dependency Injection**: [Google Wire](https://github.com/google/wire) - Compile-time DI
- **Configuration**: Environment variables with [godotenv](https://github.com/joho/godotenv)

## Prerequisites

- Go 1.25.3 or higher
- PostgreSQL database
- Wire CLI (optional, for regenerating dependency injection code)

## Getting Started

### 1. Clone the repository

```bash
git clone <repository-url>
cd todo-list-go
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Configure environment

Copy the example environment file and edit it with your database credentials:

```bash
cp .env.example .env
```

Edit `.env`:

```env
PORT=8080
DB_URL="host=localhost user=postgres password=postgres dbname=todolist port=5432 sslmode=disable"
```

### 4. Start PostgreSQL

Ensure your PostgreSQL server is running and the database exists:

```bash
createdb todolist
```

### 5. Run the application

```bash
go run cmd/*.go
```

The server will start on `http://localhost:8080` (or the port specified in your `.env`).

## API Endpoints

All endpoints are prefixed with `/api/v1`.

### Get all todos

```http
GET /api/v1/todos
```

**Response:**

```json
{
  "success": true,
  "message": "Todos recuperados",
  "data": [
    {
      "ID": 1,
      "CreatedAt": "2024-01-01T12:00:00Z",
      "UpdatedAt": "2024-01-01T12:00:00Z",
      "DeletedAt": null,
      "content": "Buy groceries",
      "status": false
    }
  ]
}
```

### Get todo by ID

```http
GET /api/v1/todos/:id
```

**Response:**

```json
{
  "success": true,
  "message": "Todo encontrado",
  "data": {
    "ID": 1,
    "CreatedAt": "2024-01-01T12:00:00Z",
    "UpdatedAt": "2024-01-01T12:00:00Z",
    "DeletedAt": null,
    "content": "Buy groceries",
    "status": false
  }
}
```

### Create todo

```http
POST /api/v1/todos
Content-Type: application/json

{
  "content": "Buy groceries",
  "status": false
}
```

**Response:**

```json
{
  "success": true,
  "message": "Todo creado con éxito",
  "data": {
    "ID": 1,
    "CreatedAt": "2024-01-01T12:00:00Z",
    "UpdatedAt": "2024-01-01T12:00:00Z",
    "DeletedAt": null,
    "content": "Buy groceries",
    "status": false
  }
}
```

### Error Response Format

```json
{
  "success": false,
  "message": "Error description",
  "error": "Detailed error message"
}
```

## Project Structure

```
.
├── cmd/                    # Application entry point
│   ├── main.go            # Main application setup and routing
│   ├── wire.go            # Wire dependency injection configuration
│   └── wire_gen.go        # Generated Wire code (do not edit)
├── config/                # Configuration management
│   └── config.go          # Environment variable loading
├── internal/              # Private application code
│   ├── domain/            # Domain entities and interfaces
│   │   └── todo.go        # Todo entity and repository interface
│   ├── repository/        # Data access layer
│   │   └── todo_repo.go   # Todo repository implementation
│   ├── service/           # Business logic layer
│   │   └── todo_service.go # Todo service with validations
│   ├── handler/           # HTTP handlers
│   │   └── todo_handler.go # Todo HTTP endpoints
│   └── db/                # Database setup
│       └── db.go          # GORM database initialization
├── pkg/                   # Public reusable packages
│   └── response/          # HTTP response utilities
│       └── response.go    # Standardized JSON responses
├── .env.example           # Example environment variables
├── go.mod                 # Go module dependencies
└── README.md             # This file
```

## Architecture

This project follows **Clean Architecture** principles:

- **Handler Layer**: Handles HTTP requests, validates input, and calls services
- **Service Layer**: Contains business logic, validations, and orchestration
- **Repository Layer**: Abstracts data access with interfaces
- **Domain Layer**: Defines core entities and repository interfaces

**Dependency Flow**: Handler → Service → Repository → Database

All dependencies are managed via **Google Wire** for compile-time dependency injection.

## Development

### Regenerating Wire Code

If you modify the dependency providers in `cmd/wire.go`, regenerate the Wire code:

```bash
go generate ./cmd
# or
wire ./cmd
```

### Adding Dependencies

```bash
go get github.com/package/name
go mod tidy
```

### Database Migrations

The application uses GORM's AutoMigrate feature. Database tables are automatically created/updated on application startup based on the domain entities.

## License

MIT
