package persist

import (
	"github.com/Masterminds/squirrel"
	"github.com/amckinney/issue-tracker/internal/model"
)

// CreateIssue creates the given issue.
func (s *sqlReadTx) CreateIssue(issue *model.Issue) (*model.Issue, error) {
	query := s.statementBuilder.Insert(
		"issues",
	).Columns(
		"entity_id",
		"id",
		"title",
		"body",
	).Values(
		issue.EntityID,
		issue.ID,
		issue.Title,
		issue.Body,
	).Suffix(
		"RETURNING *",
	)
	return s.querySingleIssue(query)
}

// UpdateIssue updates the given issue.
//
// TODO: This is an all-or-nothing update.
// We should dynamically append values
// only if they are explicitly set.
func (s *sqlReadTx) UpdateIssue(issue *model.Issue) (*model.Issue, error) {
	query := s.statementBuilder.Update(
		"issues",
	).Set(
		"title", issue.Title,
	).Set(
		"body", issue.Body,
	).Where(
		squirrel.Eq{
			"issues.id": issue.ID,
		},
	).Suffix(
		"RETURNING *",
	)
	return s.querySingleIssue(query)
}

// DeleteIssue deletes the given issue.
func (s *sqlReadTx) DeleteIssue(issueID string) (*model.Issue, error) {
	query := s.statementBuilder.Delete(
		"issues",
	).Where(
		squirrel.Eq{
			"issues.id": issueID,
		},
	).Suffix(
		"RETURNING *",
	)
	return s.querySingleIssue(query)
}
