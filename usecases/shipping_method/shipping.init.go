package shippingmethod

import "github.com/jatis/oms/repositories"

type (
	ShippingMethodUsecase struct {
		shippingMethodManager repositories.ShippingMethodManager
	}

	Opts struct {
		ShippingMethodManager repositories.ShippingMethodManager
	}
)

func New(o *Opts) *ShippingMethodUsecase {
	return &ShippingMethodUsecase{
		shippingMethodManager: o.ShippingMethodManager,
	}
}
