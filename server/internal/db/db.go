package db

import (
	"database/sql"
	"fmt"

	// Import database drivers
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/tamakiii/flashcard-app/server/internal/config"
)

// DB represents a database connection
type DB struct {
	*sql.DB
}

// Connect establishes a connection to the database
func Connect(cfg config.DatabaseConfig) (*DB, error) {
	var dsn string
	var driver string

	switch cfg.Driver {
	case "postgres":
		driver = "postgres"
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode)
	case "sqlite3":
		driver = "sqlite3"
		dsn = fmt.Sprintf("%s.db", cfg.Name)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Driver)
	}

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Create tables if they don't exist
	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return &DB{db}, nil
}

// createTables creates database tables if they don't exist
func createTables(db *sql.DB) error {
	// Create flashcards table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS flashcards (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			front TEXT NOT NULL,
			back TEXT NOT NULL,
			category TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		return err
	}

	return nil
}
