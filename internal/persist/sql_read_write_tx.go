package persist

import (
	"context"

	"github.com/Masterminds/squirrel"
)

type sqlReadWriteTx struct {
	*sqlReadTx

	statementBuilder squirrel.StatementBuilderType
}

func newSQLReadWriteTx(ctx context.Context, statementBuilder squirrel.StatementBuilderType) *sqlReadWriteTx {
	return &sqlReadWriteTx{
		sqlReadTx:        newSQLReadTx(ctx, statementBuilder),
		statementBuilder: statementBuilder,
	}
}
