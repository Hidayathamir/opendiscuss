package utils

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ITransactionManager interface {
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}

type TransactionManager struct {
	db *gorm.DB
}

func NewTransactionManager(db *gorm.DB) ITransactionManager {
	return &TransactionManager{db: db}
}

type ctxKey string

var CtxKey = ctxKey("tr")

func (tm *TransactionManager) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	tr, isHasExternalTransaction := ctx.Value(CtxKey).(*gorm.DB)

	if !isHasExternalTransaction {
		tr = tm.db.Begin()
		ctx = context.WithValue(ctx, CtxKey, tr)
	}

	err := fn(ctx)

	if !isHasExternalTransaction {
		if err != nil {
			tr.Rollback()
			return errors.Wrap(err, "error sql transaction")
		}
		tr.Commit()
	}

	return err
}
