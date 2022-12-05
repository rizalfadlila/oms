package customer

import "github.com/jatis/oms/repositories"

type (
	CustomerUsecase struct {
		customerManager repositories.CustomerManager
	}

	Opts struct {
		CustomerManager repositories.CustomerManager
	}
)

func New(o *Opts) *CustomerUsecase {
	return &CustomerUsecase{
		customerManager: o.CustomerManager,
	}
}
