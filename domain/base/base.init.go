package base

import (
	"context"
	sqlib "database/sql"
	"github.com/jatis/oms/lib/database/sql"
	"github.com/jmoiron/sqlx"
)

type (
	BaseModule struct {
		db *sql.Store
	}

	Execer interface {
		sqlx.Execer
		sqlx.ExecerContext
		NamedExec(query string, arg interface{}) (sqlib.Result, error)
		NamedExecContext(ctx context.Context, query string, arg interface{}) (sqlib.Result, error)
	}

	Queryer interface {
		sqlx.Queryer
		GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
		SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
		PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	}

	QueryExecer interface {
		Queryer
		Execer
	}
)

func New(db *sql.Store) *BaseModule {
	return &BaseModule{
		db: db,
	}
}
