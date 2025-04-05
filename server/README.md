# Flashcard App Server

This is a REST API server built with Go to support the Flashcard application.

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

## Project Structure

```
server/
├── cmd/
│   └── server/
│       └── main.go       # Application entry point
├── internal/
│   ├── api/              # API handlers and routes
│   ├── config/           # Configuration
│   ├── db/               # Database connection
│   └── models/           # Data models
├── go.mod                # Go module definition
├── go.sum                # Go module checksums
└── README.md             # This file
```
