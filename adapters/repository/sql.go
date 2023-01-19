package repository

import "context"

type Transaction interface {
	Begin(context.Context) Transaction
	Commit()
	Rollback()
	FromContext(context.Context) interface{}
	FromContextWithTrashed(context.Context) interface{}
}
