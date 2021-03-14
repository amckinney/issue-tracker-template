package issue

import (
	"github.com/amckinney/issue-tracker/internal/entity"
	"github.com/amckinney/issue-tracker/internal/persist"
	"github.com/gofrs/uuid"
)

// Controller implements the issue tracker business logic.
type Controller interface {
	CreateIssue(tx persist.ReadWriteTx, issue *entity.Issue) (*entity.Issue, error)
	GetIssue(tx persist.ReadTx, issueID uuid.UUID) (*entity.Issue, error)
	UpdateIssue(tx persist.ReadWriteTx, issue *entity.Issue) (*entity.Issue, error)
	DeleteIssue(tx persist.ReadWriteTx, issueID uuid.UUID) (*entity.Issue, error)
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
func (h *controller) CreateIssue(tx persist.ReadWriteTx, issue *entity.Issue) (*entity.Issue, error) {
	return nil, nil
}

// GetIssue gets the issue identified by the given issueID.
func (h *controller) GetIssue(tx persist.ReadTx, issueID uuid.UUID) (*entity.Issue, error) {
	return nil, nil
}

// UpdateIssue updates the given issue.
func (h *controller) UpdateIssue(tx persist.ReadWriteTx, issue *entity.Issue) (*entity.Issue, error) {
	return nil, nil
}

// DeleteIssue deletes the issue identified by the given issueID.
func (h *controller) DeleteIssue(tx persist.ReadWriteTx, issueID uuid.UUID) (*entity.Issue, error) {
	return nil, nil
}
