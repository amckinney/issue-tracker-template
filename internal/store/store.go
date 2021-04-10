package store

import "context"

// Store is a store.
type Store interface {
	Read(context.Context, func(ReadTx) error) error
	ReadWrite(context.Context, func(ReadWriteTx) error) error
	Close() error
}
