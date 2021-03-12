package issue

import (
	"context"

	"github.com/amckinney/issue-tracker/internal/entity"
	"github.com/amckinney/issue-tracker/internal/repository/issue"
	"github.com/gofrs/uuid"
)

// Controller implements the issue endpoints.
type Controller interface {
	CreateIssue(ctx context.Context, issue *entity.Issue) (*entity.Issue, error)
	GetIssue(ctx context.Context, issueID uuid.UUID) (*entity.Issue, error)
	UpdateIssue(ctx context.Context, issue *entity.Issue) (*entity.Issue, error)
	DeleteIssue(ctx context.Context, issueID uuid.UUID) (*entity.Issue, error)
}

type controller struct {
	repository issue.Repository
}

// New returns a new Controller.
func New(repository issue.Repository) Controller {
	return newController(repository)
}

func newController(repository issue.Repository) *controller {
	return &controller{
		repository: repository,
	}
}

// CreateIssue creates the given issue.
func (h *controller) CreateIssue(ctx context.Context, issue *entity.Issue) (*entity.Issue, error) {
	return nil, nil
}

// GetIssue gets the issue identified by the given issueID.
func (h *controller) GetIssue(ctx context.Context, issueID uuid.UUID) (*entity.Issue, error) {
	return nil, nil
}

// UpdateIssue updates the given issue.
func (h *controller) UpdateIssue(ctx context.Context, issue *entity.Issue) (*entity.Issue, error) {
	return nil, nil
}

// DeleteIssue deletes the issue identified by the given issueID.
func (h *controller) DeleteIssue(ctx context.Context, issueID uuid.UUID) (*entity.Issue, error) {
	return nil, nil
}
