package user

import (
	"context"

	"github.com/amckinney/issue-tracker/internal/entity"
	"github.com/gofrs/uuid"
)

// Controller implements the user endpoints.
type Controller interface {
	CreateUser(ctx context.Context, issue *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, issueID uuid.UUID) (*entity.User, error)
	UpdateUser(ctx context.Context, issue *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, issueID uuid.UUID) (*entity.User, error)
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
func (h *controller) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return nil, nil
}

// GetUser gets the user identified by the given userID.
func (h *controller) GetUser(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	return nil, nil
}

// UpdateUser updates the given user.
func (h *controller) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return nil, nil
}

// DeleteUser deletes the user identified by the given userID.
func (h *controller) DeleteUser(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	return nil, nil
}
