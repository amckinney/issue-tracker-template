package persist

import "github.com/amckinney/issue-tracker/internal/model"

// CreateIssue creates the given issue.
func (s *sqlReadTx) CreateIssue(issue *model.Issue) (*model.Issue, error) {
	return nil, nil
}

// UpdateIssue updates the given issue.
func (s *sqlReadTx) UpdateIssue(issue *model.Issue) (*model.Issue, error) {
	return nil, nil
}

// DeleteIssue deletes the given issue.
func (s *sqlReadTx) DeleteIssue(issueID string) (*model.Issue, error) {
	return nil, nil
}
