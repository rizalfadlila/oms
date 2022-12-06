package custerr

import (
	"github.com/jatis/oms/lib/log"
	"strconv"
)

const (
	ErrCodeNoContent           ErrCode = "204000"
	ErrCodeBadRequest          ErrCode = "400000"
	ErrCodeUnprocessableEntity ErrCode = "422000"
	ErrCodeUnauthorized        ErrCode = "401000"
	ErrCodeForbidden           ErrCode = "403000"
	ErrPreConditionFailed      ErrCode = "412000"
	ErrCodeNotFound            ErrCode = "404000"
	ErrCodeConflict            ErrCode = "409000"
	ErrCodeInternalError       ErrCode = "500000"
	ErrCodeTimeoutError        ErrCode = "504000"
)

type (
	ErrCode string

	ErrChain struct {
		Exception  error
		Stacktrace error
		Code       ErrCode
	}

	ErrOpt func(o *ErrChain)
)

func New(exception error, opts ...ErrOpt) *ErrChain {
	e := &ErrChain{
		Code:       ErrCode("500000"),
		Exception:  exception,
		Stacktrace: exception,
	}

	for _, o := range opts {
		o(e)
	}

	return e
}

func (err ErrChain) Error() string {
	return err.Exception.Error()
}

func WithErrCode(code ErrCode) ErrOpt {
	return func(o *ErrChain) {
		o.Code = code
	}
}

func WithStacktrace(s error) ErrOpt {
	return func(o *ErrChain) {
		o.Stacktrace = s
	}
}

func (e ErrCode) GetStatusCode() int {
	if len(e) < 3 {
		log.Errorln("status code forced to be 500 because length error code less than 3")
		return 500
	}
	s := e[0:3]
	i, _ := strconv.Atoi(string(s))
	return i
}

func As(err error) *ErrChain {
	return err.(*ErrChain)
}

func IsCustErr(err error) (bool, *ErrChain) {
	if err, ok := err.(*ErrChain); ok {
		return true, err
	}
	return false, nil
}

func ToCustErr(err error) *ErrChain {
	return New(err)
}
