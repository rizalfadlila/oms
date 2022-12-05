package order

import "github.com/jatis/oms/repositories"

type (
	OrderUsecase struct {
		orderManager repositories.OrderManager
	}

	Opts struct {
		OrderManager repositories.OrderManager
	}
)

func New(o *Opts) *OrderUsecase {
	return &OrderUsecase{
		orderManager: o.OrderManager,
	}
}
