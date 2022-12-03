package ctx

import (
	"context"
	"github.com/jmoiron/sqlx"
)

func GetSqlTx(ctx context.Context) *sqlx.Tx {
	if tx, ok := ctx.Value(keySQLTransaction).(*sqlx.Tx); ok {
		return tx
	}
	return nil
}
