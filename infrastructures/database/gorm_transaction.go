package database

import (
	"context"

	"github.com/achmadAlli/go-simple-boilerplate/adapters/repository"
	"gorm.io/gorm"
)

type gormTransaction struct {
	db *gorm.DB
}

func NewTransaction(db *gorm.DB) repository.Transaction {
	return &gormTransaction{
		db: db,
	}
}

func (trx *gormTransaction) Begin(ctx context.Context) repository.Transaction {
	return &gormTransaction{
		db: trx.db.Begin(),
	}
}

func (trx *gormTransaction) Commit() {
	trx.db.Commit()
	trx = nil
}

func (trx *gormTransaction) Rollback() {
	trx.db.Rollback()
}

func (trx *gormTransaction) FromContext(ctx context.Context, key string) interface{} {
	val := ctx.Value(key)

	if val == nil {
		return trx.db
	}

	fromContext, ok := val.(*gormTransaction)
	if !ok {
		return trx.db
	}

	return fromContext.db
}
