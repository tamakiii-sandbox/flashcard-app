package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config represents the application configuration
type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
}

// ServerConfig contains server-specific configuration
type ServerConfig struct {
	Port int `json:"port"`
}

// DatabaseConfig contains database-specific configuration
type DatabaseConfig struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	SSLMode  string `json:"sslMode"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() Config {
	return Config{
		Server: ServerConfig{
			Port: 8080,
		},
		Database: DatabaseConfig{
			Driver:   "sqlite3",
			Host:     "localhost",
			Port:     5432,
			Name:     "flashcards",
			User:     "postgres",
			Password: "postgres",
			SSLMode:  "disable",
		},
	}
}

// Load loads the configuration from a file or environment variables
func Load() (Config, error) {
	cfg := DefaultConfig()

	// Check if config file exists
	configPath := "config.json"
	if configFile := os.Getenv("CONFIG_FILE"); configFile != "" {
		configPath = configFile
	}

	if _, err := os.Stat(configPath); err == nil {
		file, err := os.Open(configPath)
		if err != nil {
			return cfg, err
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&cfg); err != nil {
			return cfg, err
		}
	}

	// Override with environment variables if they exist
	if port := os.Getenv("SERVER_PORT"); port != "" {
		var serverPort int
		if _, err := fmt.Sscanf(port, "%d", &serverPort); err == nil {
			cfg.Server.Port = serverPort
		}
	}

	if dbDriver := os.Getenv("DB_DRIVER"); dbDriver != "" {
		cfg.Database.Driver = dbDriver
	}
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		cfg.Database.Host = dbHost
	}
	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		var port int
		if _, err := fmt.Sscanf(dbPort, "%d", &port); err == nil {
			cfg.Database.Port = port
		}
	}
	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		cfg.Database.Name = dbName
	}
	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		cfg.Database.User = dbUser
	}
	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		cfg.Database.Password = dbPassword
	}
	if dbSSLMode := os.Getenv("DB_SSL_MODE"); dbSSLMode != "" {
		cfg.Database.SSLMode = dbSSLMode
	}

	return cfg, nil
}