package models

import (
	"github.com/jatis/oms/lib/datatype"
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
