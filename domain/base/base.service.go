package base

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	libctx "github.com/jatis/oms/lib/ctx"
	"github.com/jatis/oms/lib/log"
	"github.com/jatis/oms/repositories"
	"github.com/jmoiron/sqlx"
)

type SqlQueryOperator string

var (
	SqlAndOperator SqlQueryOperator = " and "
)

func (b *BaseModule) WithTransaction(ctx context.Context, fn repositories.TransactionFunc) error {
	if parentTx := libctx.GetSqlTx(ctx); parentTx != nil {
		return fn(ctx)
	}

	tx, err := b.db.GetMaster().BeginTxx(ctx, nil)

	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func(tx *sqlx.Tx) {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.WithError(err).Errorln("failed on rollback transaction")
		}
	}(tx)

	if err := fn(libctx.SetSqlTx(ctx, tx)); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (b *BaseModule) GenerateUUID() int64 {
	// TODO : implement snowflake
	return 0
}

func (b *BaseModule) GetQueryerExecerFromContext(ctx context.Context) QueryExecer {
	if tx := libctx.GetSqlTx(ctx); tx != nil {
		return tx
	}

	return b.db.GetMaster()
}

func (b *BaseModule) GetQueryerFromContext(ctx context.Context) Queryer {
	if tx := libctx.GetSqlTx(ctx); tx != nil {
		return tx
	}

	return b.db.GetMaster()
}

func (b *BaseModule) GetExecerFromContext(ctx context.Context) Execer {
	if tx := libctx.GetSqlTx(ctx); tx != nil {
		return tx
	}

	return b.db.GetMaster()
}

func (b *BaseModule) GetQueryer() Queryer {
	return b.db.GetMaster()
}
