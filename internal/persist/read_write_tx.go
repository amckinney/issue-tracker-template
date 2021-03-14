package persist

import "github.com/amckinney/issue-tracker/internal/model"

// ReadWriteTx provides a read-write transaction.
type ReadWriteTx interface {
	ReadTx

	CreateIssue(issue *model.Issue) (*model.Issue, error)
	UpdateIssue(issue *model.Issue) (*model.Issue, error)
	DeleteIssue(issueID string) (*model.Issue, error)
}
