package user

import (
	"context"

	"github.com/amckinney/issue-tracker/internal/entity"
	"github.com/gofrs/uuid"
)

// Controller implements the user endpoints.
type Controller interface {
	CreateUser(ctx context.Context, user *entity.User)
	GetUser(ctx context.Context, userID uuid.UUID)
	UpdateUser(ctx context.Context, user *entity.User)
	DeleteUser(ctx context.Context, userID uuid.UUID)
}

type controller struct{}

// New returns a new Controller.
func New() Controller {
	return newController()
}

func newController() *controller {
	return &controller{}
}

// CreateUser creates the given user.
func (h *controller) CreateUser(ctx context.Context, user *entity.User) {
	// TODO
	return
}

// GetUser gets the user identified by the given userID.
func (h *controller) GetUser(ctx context.Context, userID uuid.UUID) {
	// TODO
	return
}

// UpdateUser updates the given user.
func (h *controller) UpdateUser(ctx context.Context, user *entity.User) {
	// TODO
	return
}

// DeleteUser deletes the user identified by the given userID.
func (h *controller) DeleteUser(ctx context.Context, userID uuid.UUID) {
	// TODO
	return
}
