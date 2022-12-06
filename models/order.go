package models

import (
	"github.com/jatis/oms/lib/datatype"
	"github.com/jatis/oms/lib/util"
)

type Order struct {
	ID                  int64                  `db:"id" json:"id"`
	CustomerID          int64                  `db:"customer_id" json:"customer_id"`
	EmployeeID          int64                  `db:"employee_id" json:"employee_id"`
	PurchaseOrderNumber string                 `db:"purchase_order_number" json:"purchase_order_number"`
	OrderDate           datatype.DateInJakarta `db:"order_date" json:"order_date"`
	ShipDate            datatype.DateInJakarta `db:"ship_date" json:"ship_date"`
	ShippingMethodID    int64                  `db:"shipping_method_id" json:"shipping_method_id"`
	FreightCharge       float64                `db:"freight_charge" json:"freight_charge"`
	Taxes               float64                `db:"taxes" json:"taxes"`
	PaymentReceived     int16                  `db:"payment_received" json:"payment_received"`
	Comment             string                 `db:"comment" json:"comment"`
}

type OrderDetail struct {
	ID        int64   `db:"id" json:"id"`
	OrderID   int64   `db:"order_id" json:"order_id"`
	ProductID int64   `db:"product_id" json:"product_id"`
	Quantity  int     `db:"quantity" json:"quantity"`
	UnitPrice float64 `db:"unit_price" json:"unit_price"`
	Discount  float64 `db:"discount" json:"discount"`
	SubTotal  float64 `db:"-" json:"sub_total"`
}

func NewOrderFromRowCSV(data []interface{}) *Order {
	order := Order{
		PurchaseOrderNumber: util.InterfaceToString(data[2]),
		OrderDate:           datatype.DateInJakarta(util.InterfaceToString(data[3])),
		ShipDate:            datatype.DateInJakarta(util.InterfaceToString(data[4])),
		FreightCharge:       util.InterfaceToFloat64(data[6]),
		Taxes:               util.InterfaceToFloat64(data[7]),
		PaymentReceived:     int16(util.InterfaceToInt(data[8])),
		Comment:             util.InterfaceToString(data[9]),
	}

	return &order
}

func NewOrderDetailFromRowCSV(data []interface{}) *OrderDetail {
	orderDetail := OrderDetail{
		Quantity: int(util.InterfaceToInt(data[2])),
		Discount: util.InterfaceToFloat64(data[3]),
	}

	return &orderDetail
}
