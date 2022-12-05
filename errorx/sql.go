package errorx

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jatis/oms/constants"
	"github.com/jatis/oms/lib/custerr"
)

type SqlErrorType string

var (
	SqlQuery       SqlErrorType = "sql-query"
	SqlTransaction SqlErrorType = "sql-transaction"

	sqlErr = map[SqlErrorType]error{
		SqlQuery:       constants.ErrQueryRequest,
		SqlTransaction: constants.ErrTransactionalRequest,
	}
)

func SqlNoRows() error {
	return custerr.New(fmt.Errorf("%w: there were no rows", constants.ErrNoData),
		custerr.WithErrCode(custerr.ErrCodeNoContent),
	)
}

func SqlError(err error, typ SqlErrorType) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return SqlNoRows()
	}

	return custerr.New(sqlErr[typ], custerr.WithStacktrace(err))
}
