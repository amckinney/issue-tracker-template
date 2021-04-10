package store

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/amckinney/issue-tracker/internal/model"
)

// CreateIssue creates the given issue.
func (s *sqlReadTx) CreateIssue(issue *model.Issue) (*model.Issue, error) {
	insertQuery := s.statementBuilder.Insert(
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
	)
	result, err := insertQuery.ExecContext(s.ctx)
	if err != nil {
		return nil, err
	}
	if affected, err := result.RowsAffected(); err != nil || affected == 0 {
		return nil, fmt.Errorf("expected exactly one row to be affected, but had %d affected", affected)
	}
	selectQuery := s.statementBuilder.Select(
		"*",
	).From(
		"issues",
	).Where(
		squirrel.Eq{
			"id": issue.ID,
		},
	)
	return s.querySingleIssue(selectQuery)
}

// UpdateIssue updates the given issue.
func (s *sqlReadTx) UpdateIssue(issue *model.Issue) (*model.Issue, error) {
	updateQuery := s.statementBuilder.Update(
		"issues",
	).Set(
		"title", issue.Title,
	).Set(
		"body", issue.Body,
	).Where(
		squirrel.Eq{
			"id": issue.ID,
		},
	)
	if _, err := updateQuery.QueryContext(s.ctx); err != nil {
		return nil, err
	}
	selectQuery := s.statementBuilder.Select(
		"*",
	).From(
		"issues",
	).Where(
		squirrel.Eq{
			"id": issue.ID,
		},
	)
	return s.querySingleIssue(selectQuery)
}

// DeleteIssue deletes the given issue.
func (s *sqlReadTx) DeleteIssue(issueID string) (*model.Issue, error) {
	selectQuery := s.statementBuilder.Select(
		"*",
	).From(
		"issues",
	).Where(
		squirrel.Eq{
			"id": issueID,
		},
	)
	issue, err := s.querySingleIssue(selectQuery)
	if err != nil {
		return nil, err
	}
	deleteQuery := s.statementBuilder.Delete(
		"issues",
	).Where(
		squirrel.Eq{
			"id": issueID,
		},
	)
	if _, err := deleteQuery.QueryContext(s.ctx); err != nil {
		return nil, err
	}
	return issue, nil
}
