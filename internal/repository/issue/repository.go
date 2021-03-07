package issue

import (
	"github.com/amckinney/issue-tracker/internal/entity"
	"github.com/amckinney/issue-tracker/internal/persist"
	"github.com/gofrs/uuid"
)

// Repository provides a persistence abstraction for issues.
type Repository interface {
	CreateIssue(tx persist.ReadWriteTx, issue *entity.Issue) (*entity.Issue, error)
	GetIssue(tx persist.ReadTx, issueID uuid.UUID) (*entity.Issue, error)
	UpdateIssue(tx persist.ReadWriteTx, issue *entity.Issue) (*entity.Issue, error)
	DeleteIssue(tx persist.ReadWriteTx, issueID uuid.UUID) (*entity.Issue, error)
}

type repository struct{}

// New returns a new Repository.
func New() Repository {
	return newRepository()
}

func newRepository() *repository {
	return &repository{}
}

// CreateIssue creates the given issue.
func (h *repository) CreateIssue(tx persist.ReadWriteTx, issue *entity.Issue) (*entity.Issue, error) {
	return nil, nil
}

// GetIssue gets the issue identified by the given issueID.
func (h *repository) GetIssue(tx persist.ReadTx, issueID uuid.UUID) (*entity.Issue, error) {
	return nil, nil
}

// UpdateIssue updates the given issue.
func (h *repository) UpdateIssue(tx persist.ReadWriteTx, issue *entity.Issue) (*entity.Issue, error) {
	return nil, nil
}

// DeleteIssue deletes the issue identified by the given issueID.
func (h *repository) DeleteIssue(tx persist.ReadWriteTx, issueID uuid.UUID) (*entity.Issue, error) {
	return nil, nil
}
