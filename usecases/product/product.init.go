package product

import "github.com/jatis/oms/repositories"

type (
	ProductUsecase struct {
		productManager repositories.ProductManager
	}

	Opts struct {
		ProductManager repositories.ProductManager
	}
)

func New(o *Opts) *ProductUsecase {
	return &ProductUsecase{
		productManager: o.ProductManager,
	}
}
