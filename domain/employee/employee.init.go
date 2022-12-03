package employee

import (
	"github.com/jatis/oms/domain/base"
	"github.com/jatis/oms/lib/database/sql"
	"github.com/jatis/oms/repositories"
)

type (
	module struct {
		*base.BaseModule
		db *sql.Store
	}

	Opts struct {
		DB *sql.Store
	}
)

func New(o *Opts) repositories.EmployeeManager {
	return &module{
		db:         o.DB,
		BaseModule: base.New(o.DB),
	}
}
