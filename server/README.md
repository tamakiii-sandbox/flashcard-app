# Flashcard App Server

This is a REST API server built with Go to support the Flashcard application. The API is defined and validated using OpenAPI 3.0 specifications.

## Setup

### Prerequisites

- Go 1.20 or later
- PostgreSQL (optional, can use SQLite for development)

### Installation

1. Install Go dependencies:
   ```
   go mod download
   ```

2. Run the server:
   ```
   go run cmd/server/main.go
   ```

## API Endpoints

- `GET /api/flashcards` - Get all flashcards
- `GET /api/flashcards/:id` - Get a specific flashcard
- `POST /api/flashcards` - Create a new flashcard
- `PUT /api/flashcards/:id` - Update a flashcard
- `DELETE /api/flashcards/:id` - Delete a flashcard

All endpoints are defined in the OpenAPI specification file at `api/openapi.yaml`.

## OpenAPI Integration

This project uses OpenAPI 3.0 for API definition, validation, and documentation. Key features include:

- **API Specification**: The complete API is defined in `api/openapi.yaml`
- **Code Generation**: Server code is generated from the OpenAPI spec using `oapi-codegen`
- **Request Validation**: Incoming requests are automatically validated against the specification
- **Type Safety**: Generated Go types match the API specification

### Updating the API

When making changes to the API:

1. Update the OpenAPI specification in `api/openapi.yaml`
2. Run the code generation script:
   ```
   ./generate.sh
   ```
3. Implement any new endpoints in `internal/api/server_impl.go`

## Project Structure

```
server/
├── api/
│   └── openapi.yaml     # OpenAPI specification file
├── cmd/
│   └── server/
│       └── main.go      # Application entry point
├── internal/
│   ├── api/
│   │   ├── api.gen.go   # Generated code from OpenAPI spec
│   │   ├── router.go    # API router with OpenAPI validation
│   │   └── server_impl.go # API implementation
│   ├── config/          # Configuration
│   ├── db/              # Database connection
│   └── models/          # Data models
├── generate.sh          # Script to generate code from OpenAPI spec
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
└── README.md            # This file
```

## Benefits of OpenAPI

- **Documentation**: API is well-documented in a standard format
- **Validation**: Automatic request/response validation
- **Client Generation**: Can generate client SDKs for various languages
- **Testing**: Compatible with tools that support OpenAPI testing
