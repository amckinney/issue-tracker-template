package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"go.uber.org/multierr"
)

type sqlStore struct {
	db               *sql.DB
	statementBuilder squirrel.StatementBuilderType
}

// NewSQL returns a new store.DB backed by *sql.DB.
func NewSQL(db *sql.DB) (Store, error) {
	return &sqlStore{
		db:               db,
		statementBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(db),
	}, nil
}

func (s *sqlStore) Read(ctx context.Context, fn func(ReadTx) error) (retErr error) {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{
		ReadOnly: true,
	})
	if err != nil {
		return fmt.Errorf("failed to create read transaction: %w", err)
	}
	defer func() {
		retErr = multierr.Append(retErr, tx.Rollback())
	}()

	readTx := newSQLReadTx(ctx, s.statementBuilder.RunWith(tx))
	if err := fn(readTx); err != nil {
		return err
	}

	return nil
}

func (s *sqlStore) ReadWrite(ctx context.Context, fn func(ReadWriteTx) error) (retErr error) {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{
		ReadOnly: false,
	})
	if err != nil {
		return fmt.Errorf("failed to create read-write transaction: %w", err)
	}
	defer func() {
		if retErr != nil {
			retErr = multierr.Append(retErr, tx.Rollback())
		}
	}()

	readWriteTx := newSQLReadWriteTx(ctx, s.statementBuilder.RunWith(tx))
	if err := fn(readWriteTx); err != nil {
		return err
	}

	return tx.Commit()
}

func (s *sqlStore) Close() error {
	return s.db.Close()
}
