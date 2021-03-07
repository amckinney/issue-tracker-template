package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

// User represents a user.
type User struct {
	ID        uuid.UUID
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
