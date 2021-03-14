package model

import (
	"time"
)

// Issue represents an issue.
type Issue struct {
	EntityID  string    `json:"entity_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
}
