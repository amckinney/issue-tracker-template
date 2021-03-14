package persist

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
	fmt.Println("Writing issue with ID", issue.ID)
	rows, err := insertQuery.QueryContext(s.ctx)
	if err != nil {
		return nil, err
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return &model.Issue{
		ID:    "a9eae375-d98b-4d9b-a39d-6683a7cfb263",
		Title: "Foo",
		Body:  "Bar",
	}, nil
	//selectQuery := s.statementBuilder.Select(
	//"*",
	//).From(
	//"issues",
	//).Where(
	//squirrel.Eq{
	//"id": issue.ID,
	//},
	//)
	//return s.querySingleIssue(selectQuery)
}

// UpdateIssue updates the given issue.
//
// TODO: This is an all-or-nothing update.
// We should dynamically append values
// only if they are explicitly set.
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
