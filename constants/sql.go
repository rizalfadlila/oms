package constants

import "errors"

var (
	ErrNoData               = errors.New("no data found")
	ErrQueryRequest         = errors.New("error on performing query request")
	ErrTransactionalRequest = errors.New("error on performing transactional request")
)
