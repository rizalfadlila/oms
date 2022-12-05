package employee

import "github.com/jatis/oms/repositories"

type (
	EmployeeUsecase struct {
		employeeManager repositories.EmployeeManager
	}

	Opts struct {
		EmployeeManager repositories.EmployeeManager
	}
)

func New(o *Opts) *EmployeeUsecase {
	return &EmployeeUsecase{
		employeeManager: o.EmployeeManager,
	}
}
