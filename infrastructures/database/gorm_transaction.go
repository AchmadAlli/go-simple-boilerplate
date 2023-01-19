package database

import (
	"context"

	"github.com/achmadAlli/go-simple-boilerplate/adapters/repository"
	"gorm.io/gorm"
)

type gormTransaction struct {
	ctx context.Context
	db  *gorm.DB
}

func NewTransaction() repository.Transaction {
	return &gormTransaction{}
}

func (trx *gormTransaction) Begin(ctx context.Context) repository.Transaction {
	return &gormTransaction{
		ctx: ctx,
		db:  trx.db.Begin(),
	}
}

func (trx *gormTransaction) Commit() {
	trx.db.Commit()
	trx = nil
}

func (trx *gormTransaction) Rollback() {
	trx.db.Rollback()
}

func (trx *gormTransaction) FromContext(ctx context.Context) interface{} {
	val := ctx.Value(trx.ctx)
	if val == nil {
		return trx.db
	}

	fromContext, ok := val.(*gormTransaction)
	if !ok {
		return trx.db
	}

	return fromContext.db
}

func (trx *gormTransaction) FromContextWithTrashed(ctx context.Context) interface{} {
	val := ctx.Value(trx.ctx)
	if val == nil {
		return trx.db.Unscoped()
	}

	fromContext, ok := val.(*gormTransaction)
	if !ok {
		return trx.db.Unscoped()
	}

	return fromContext.db.Unscoped()
}
