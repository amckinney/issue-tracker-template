package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

// Issue represents an issue.
type Issue struct {
	ID        uuid.UUID
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
