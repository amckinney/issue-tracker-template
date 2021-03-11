package persist

import (
	"context"

	"github.com/Masterminds/squirrel"
)

type sqlReadTx struct {
	ctx              context.Context
	statementBuilder squirrel.StatementBuilderType
}

func newSQLReadTx(ctx context.Context, statementBuilder squirrel.StatementBuilderType) *sqlReadTx {
	return &sqlReadTx{
		ctx:              ctx,
		statementBuilder: statementBuilder,
	}
}

// Context returns this transaction's context.
func (s *sqlReadTx) Context() context.Context {
	return s.ctx
}
