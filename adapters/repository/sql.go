package repository

import "context"

type Transaction interface {
	Begin(context.Context) Transaction
	Commit()
	Rollback()
	FromContext(ctx context.Context, key string) interface{}
}
