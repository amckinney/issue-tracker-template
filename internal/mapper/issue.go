package mapper

import (
	"errors"
	"fmt"

	"github.com/amckinney/issue-tracker/api"
	"github.com/amckinney/issue-tracker/internal/entity"
	"github.com/gofrs/uuid"
)

// IssueFromAPI maps the given *api.Issue into an
// *entity.Issue.
func IssueFromAPI(issue *api.Issue) (*entity.Issue, error) {
	if issue == nil {
		return nil, errors.New("received a nil issue")
	}
	issueID, err := uuid.FromString(issue.ID)
	if err != nil {
		return nil, fmt.Errorf("id %q is not a valid uuid: %w", issue.ID, err)
	}
	return &entity.Issue{
		ID:    issueID,
		Title: issue.Title,
		Body:  issue.Body,
	}, nil
}

// IssueToAPI maps the given *entity.Issue into an
// *api.Issue.
func IssueToAPI(issue *entity.Issue) (*api.Issue, error) {
	if issue == nil {
		return nil, errors.New("received a nil issue")
	}
	return &api.Issue{
		ID:        issue.ID.String(),
		Title:     issue.Title,
		Body:      issue.Body,
		CreatedAt: issue.CreatedAt,
		UpdatedAt: issue.UpdatedAt,
	}, nil
}
