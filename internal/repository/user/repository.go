package user

import (
	"github.com/amckinney/issue-tracker/internal/entity"
	"github.com/amckinney/issue-tracker/internal/persist"
	"github.com/gofrs/uuid"
)

// Repository provides a persistence abstraction for issues.
type Repository interface {
	CreateUser(tx persist.ReadWriteTx, issue *entity.User) (*entity.User, error)
	GetUser(tx persist.ReadTx, issueID uuid.UUID) (*entity.User, error)
	UpdateUser(tx persist.ReadWriteTx, issue *entity.User) (*entity.User, error)
	DeleteUser(tx persist.ReadWriteTx, issueID uuid.UUID) (*entity.User, error)
}

type repository struct{}

// New returns a new Repository.
func New() Repository {
	return newRepository()
}

func newRepository() *repository {
	return &repository{}
}

// CreateUser creates the given issue.
func (h *repository) CreateUser(tx persist.ReadWriteTx, issue *entity.User) (*entity.User, error) {
	return nil, nil
}

// GetUser gets the issue identified by the given issueID.
func (h *repository) GetUser(tx persist.ReadTx, issueID uuid.UUID) (*entity.User, error) {
	return nil, nil
}

// UpdateUser updates the given issue.
func (h *repository) UpdateUser(tx persist.ReadWriteTx, issue *entity.User) (*entity.User, error) {
	return nil, nil
}

// DeleteUser deletes the issue identified by the given issueID.
func (h *repository) DeleteUser(tx persist.ReadWriteTx, issueID uuid.UUID) (*entity.User, error) {
	return nil, nil
}
