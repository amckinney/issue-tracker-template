package persist

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/amckinney/issue-tracker/internal/model"
	"go.uber.org/multierr"
)

// GetIssue gets the issue identified by the given issueID.
func (s *sqlReadTx) GetIssue(issueID string) (*model.Issue, error) {
	query := s.statementBuilder.Select(
		"*",
	).From(
		"issues",
	).Where(
		squirrel.Eq{
			"issues.id": issueID,
		},
	)
	return s.querySingleIssue(query)
}

// querier is used to generalize a variety of query commands
// under a common interface.
type querier interface {
	QueryContext(context.Context) (*sql.Rows, error)
}

// querySingleIssue executes the given querier and returns the
// associated issue.
func (s *sqlReadTx) querySingleIssue(querier querier) (*model.Issue, error) {
	rows, err := querier.QueryContext(s.ctx)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, multierr.Append(
			fmt.Errorf("issue querier returned no results"),
			rows.Close(),
		)
	}
	var issue model.Issue
	if err := rows.Scan(
		&issue.EntityID,
		&issue.CreatedAt,
		&issue.UpdatedAt,
		&issue.ID,
		&issue.Title,
		&issue.Body,
	); err != nil {
		return nil, multierr.Append(err, rows.Close())
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &issue, nil
}
