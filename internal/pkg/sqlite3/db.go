package sqlite3

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

// New constructs a new database backed by SQLite3.
func New(config *Config) (*sql.DB, error) {
	if config == nil {
		return nil, errors.New("received a nil sqlite3 configuration")
	}
	if config.Source == "" {
		return nil, fmt.Errorf("sqlite3 requires a source")
	}
	return sql.Open("sqlite3", config.Source)
}
