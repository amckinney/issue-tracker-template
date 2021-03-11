package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
)

// Config represents the configuration used to initialize a
// SQLite3 database.
type Config struct {
	Source string
}

// New constructs a new *sql.DB backed by SQLite3.
func New(config *Config) (*sql.DB, error) {
	if config == nil {
		return nil, errors.New("received a nil sqlite3 configuration")
	}
	if config.Source == "" {
		return nil, fmt.Errorf("sqlite3 requires a source")
	}
	db, err := sql.Open("sqlite3", config.Source)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite3 database: %w", err)
	}
	return db, db.Ping()
}
