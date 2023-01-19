package repository

import "context"

type SQL interface {
	Execute(ctx context.Context, queryString string, result ...interface{}) error
	Query(ctx context.Context, queryString string, result ...interface{}) (Rows, error)
	QueryRow(ctx context.Context, queryString string, result ...interface{}) Row
	BeginTx(ctx context.Context) (Transaction, error)
}

type Transaction interface {
	Execute(ctx context.Context, queryString string, result ...interface{}) error
	Query(ctx context.Context, queryString string, result ...interface{}) (Rows, error)
	QueryRow(ctx context.Context, queryString string, result ...interface{}) Row
	Commit() error
	Rollback() error
}

type Rows interface {
	Scan(dest ...interface{}) error
	Next() bool
	Err() error
	Close() error
}

type Row interface {
	Scan(dest ...interface{}) error
}
