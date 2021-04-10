package issue

import (
	"github.com/amckinney/issue-tracker/internal/entity"
	"github.com/amckinney/issue-tracker/internal/mapper"
	"github.com/amckinney/issue-tracker/internal/store"
	"github.com/gofrs/uuid"
)

// Controller implements the issue tracker business logic.
type Controller interface {
	CreateIssue(tx store.ReadWriteTx, issue *entity.Issue) (*entity.Issue, error)
	GetIssue(tx store.ReadTx, issueID uuid.UUID) (*entity.Issue, error)
	UpdateIssue(tx store.ReadWriteTx, issue *entity.Issue) (*entity.Issue, error)
	DeleteIssue(tx store.ReadWriteTx, issueID uuid.UUID) (*entity.Issue, error)
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
func (h *controller) CreateIssue(tx store.ReadWriteTx, issue *entity.Issue) (*entity.Issue, error) {
	model, err := mapper.IssueToModel(issue)
	if err != nil {
		return nil, err
	}
	entityID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	model.EntityID = entityID.String()
	model.ID = id.String()
	createdIssue, err := tx.CreateIssue(model)
	if err != nil {
		return nil, err
	}
	return mapper.IssueFromModel(createdIssue)
}

// GetIssue gets the issue identified by the given issueID.
func (h *controller) GetIssue(tx store.ReadTx, issueID uuid.UUID) (*entity.Issue, error) {
	gotIssue, err := tx.GetIssue(issueID.String())
	if err != nil {
		return nil, err
	}
	return mapper.IssueFromModel(gotIssue)
}

// UpdateIssue updates the given issue.
func (h *controller) UpdateIssue(tx store.ReadWriteTx, issue *entity.Issue) (*entity.Issue, error) {
	model, err := mapper.IssueToModel(issue)
	if err != nil {
		return nil, err
	}
	updatedIssue, err := tx.UpdateIssue(model)
	if err != nil {
		return nil, err
	}
	return mapper.IssueFromModel(updatedIssue)
}

// DeleteIssue deletes the issue identified by the given issueID.
func (h *controller) DeleteIssue(tx store.ReadWriteTx, issueID uuid.UUID) (*entity.Issue, error) {
	deletedIssue, err := tx.DeleteIssue(issueID.String())
	if err != nil {
		return nil, err
	}
	return mapper.IssueFromModel(deletedIssue)
}
