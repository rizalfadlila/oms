package initiator

import (
	"github.com/jatis/oms/init/service"
	"github.com/jatis/oms/usecases/customer"
	"github.com/jatis/oms/usecases/employee"
	"github.com/jatis/oms/usecases/order"
	"github.com/jatis/oms/usecases/product"
	shippingmethod "github.com/jatis/oms/usecases/shipping_method"
)

func (i *initiator) newUsecase() {
	i.usecase = &service.Usecases{
		CustomerUsecase:       i.NewCustomerUsecase(),
		EmployeeUsecase:       i.NewEmployeeUsecase(),
		ProductUsecase:        i.NewProductUsecase(),
		ShippingMethodUsecase: i.NewShippingMethodUsecase(),
		OrderUsecase:          i.NewOrderUsecase(),
	}
}

func (i *initiator) NewCustomerUsecase() *customer.CustomerUsecase {
	opt := &customer.Opts{CustomerManager: i.NewCustomerManager()}
	return customer.New(opt)
}

func (i *initiator) NewEmployeeUsecase() *employee.EmployeeUsecase {
	opt := &employee.Opts{EmployeeManager: i.NewEmployeeManager()}
	return employee.New(opt)
}

func (i *initiator) NewProductUsecase() *product.ProductUsecase {
	opt := &product.Opts{ProductManager: i.NewProductManager()}
	return product.New(opt)
}

func (i *initiator) NewShippingMethodUsecase() *shippingmethod.ShippingMethodUsecase {
	opt := &shippingmethod.Opts{ShippingMethodManager: i.NewShippingMethodManager()}
	return shippingmethod.New(opt)
}

func (i *initiator) NewOrderUsecase() *order.OrderUsecase {
	opt := &order.Opts{OrderManager: i.NewOrderManager()}
	return order.New(opt)
}
