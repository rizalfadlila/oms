package ctx

import (
	"context"
)

const (
	keySQLTransaction = "sql-transaction"
	keyTokenClaim     = "token-token_claim"
)

func SetSqlTx(ctx context.Context, value interface{}) context.Context {
	return context.WithValue(ctx, keySQLTransaction, value)
}
