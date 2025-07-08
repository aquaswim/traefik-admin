# Traefik Admin

A simple Go Fiber application using hexagonal architecture.

## Project Structure

This project follows the standard Go project layout with a hexagonal architecture pattern:

- `cmd/traefik-admin/`: Application entry point
- `internal/`: Private application code
  - `domain/`: Core business logic
  - `application/`: Use cases
  - `adapters/`: Infrastructure adapters (HTTP, database, etc.)

## Getting Started

### Prerequisites

- Go 1.24 or higher

### Running the Application

```bash
# Download dependencies
go mod tidy

# Run the application
go run cmd/traefik-admin/main.go
```

The application will start on port 3000. You can access it at http://localhost:3000.

## API Endpoints

- `GET /api`: Returns a "Hello, World!" message