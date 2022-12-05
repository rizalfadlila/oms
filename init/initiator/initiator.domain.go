package initiator

import (
	"github.com/jatis/oms/domain/customer"
	"github.com/jatis/oms/domain/employee"
	"github.com/jatis/oms/domain/order"
	"github.com/jatis/oms/domain/product"
	shippingmethod "github.com/jatis/oms/domain/shipping_method"
	"github.com/jatis/oms/init/service"
	"github.com/jatis/oms/repositories"
)

func (i *initiator) newDomain() {
	i.domain = &service.Domains{
		CustomerManager:       i.NewCustomerManager(),
		EmployeeManager:       i.NewEmployeeManager(),
		ShippingMethodManager: i.NewShippingMethodManager(),
		ProductManager:        i.NewProductManager(),
		OrderManager:          i.NewOrderManager(),
	}
}

func (i *initiator) NewCustomerManager() repositories.CustomerManager {
	opt := &customer.Opts{DB: i.basic.MariaClient}
	return customer.New(opt)
}

func (i *initiator) NewEmployeeManager() repositories.EmployeeManager {
	opt := &employee.Opts{DB: i.basic.MariaClient}
	return employee.New(opt)
}

func (i *initiator) NewProductManager() repositories.ProductManager {
	opt := &product.Opts{DB: i.basic.MariaClient}
	return product.New(opt)
}

func (i *initiator) NewOrderManager() repositories.OrderManager {
	opt := &order.Opts{DB: i.basic.MariaClient}
	return order.New(opt)
}

func (i *initiator) NewShippingMethodManager() repositories.ShippingMethodManager {
	opt := &shippingmethod.Opts{DB: i.basic.MariaClient}
	return shippingmethod.New(opt)
}
