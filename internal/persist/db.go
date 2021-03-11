package persist

import "context"

// DB is a database.
type DB interface {
	// Read starts a read-only transaction.
	Read(context.Context, func(ReadTx) error) error
	// ReadWrite starts a read-write transaction. Results are only
	// committed if the given function does not return an error.
	ReadWrite(context.Context, func(ReadWriteTx) error) error
	// Close closes the database.
	Close() error
}
