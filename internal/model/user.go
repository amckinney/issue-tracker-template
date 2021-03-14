package model

import "time"

// User represents a user.
type User struct {
	EntityID  string    `json:"entity_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        string    `json:"id"`
	Username  string    `json:"username"`
}
