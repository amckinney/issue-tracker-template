package persist

import "github.com/amckinney/issue-tracker/internal/model"

// ReadTx provides a read-only transaction.
type ReadTx interface {
	GetIssue(issueID string) (*model.Issue, error)
}
