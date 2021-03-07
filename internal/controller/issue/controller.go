package issue

import (
	"context"

	"github.com/amckinney/issue-tracker/internal/entity"
	"github.com/gofrs/uuid"
)

// Controller implements the issue endpoints.
type Controller interface {
	CreateIssue(ctx context.Context, issue *entity.Issue)
	GetIssue(ctx context.Context, issueID uuid.UUID)
	UpdateIssue(ctx context.Context, issue *entity.Issue)
	DeleteIssue(ctx context.Context, issueID uuid.UUID)
}

type controller struct{}

// New returns a new Controller.
func New() Controller {
	return newController()
}

func newController() *controller {
	return &controller{}
}

// CreateIssue creates the given issue.
func (h *controller) CreateIssue(ctx context.Context, issue *entity.Issue) {
	// TODO
	return
}

// GetIssue gets the issue identified by the given issueID.
func (h *controller) GetIssue(ctx context.Context, issueID uuid.UUID) {
	// TODO
	return
}

// UpdateIssue updates the given issue.
func (h *controller) UpdateIssue(ctx context.Context, issue *entity.Issue) {
	// TODO
	return
}

// DeleteIssue deletes the issue identified by the given issueID.
func (h *controller) DeleteIssue(ctx context.Context, issueID uuid.UUID) {
	// TODO
	return
}
