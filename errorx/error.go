package errorx

import (
	"errors"
	"github.com/jatis/oms/constants"
	"github.com/jatis/oms/lib/custerr"
)

func IsErrorNoData(err error) bool {
	if ok, err := custerr.IsCustErr(err); ok && errors.Is(err.Exception, constants.ErrNoData) {
		return true
	}
	return false
}
