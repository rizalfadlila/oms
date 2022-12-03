package models

type ReportOrder struct {
	CustomerName       string `db:"customer_name"`
	EmployeeName       string `db:"employee_name"`
	ShippingMethodName string `db:"shipping_method_name"`
}

type ReportOrderDetail struct {
	ID        int64   `db:"product_id" json:"product_id"`
	OrderID   int64   `db:"product_name" json:"product_name"`
	Quantity  int     `db:"quantity" json:"quantity"`
	UnitPrice float64 `db:"unit_price" json:"unit_price"`
	Discount  float64 `db:"discount" json:"discount"`
	SubTotal  float64 `db:"-" json:"sub_total"`
}

type ResponseReportOrder struct {
	CustomerName       string              `db:"customer_name" json:"customer_name"`
	EmployeeName       string              `db:"employee_name" json:"employee_name"`
	ShippingMethodName string              `db:"shipping_method_name" json:"shipping_method_name"`
	Items              []ReportOrderDetail `db:"-" json:"items"`
	TotalPayment       float64             `db:"-" json:"total_payment"`
}

func ComposeReportData(order ReportOrder, detail []ReportOrderDetail) *ResponseReportOrder {
	data := &ResponseReportOrder{
		CustomerName:       order.CustomerName,
		EmployeeName:       order.CustomerName,
		ShippingMethodName: order.ShippingMethodName,
	}

	for _, item := range detail {
		item.SubTotal = item.UnitPrice - item.Discount
		data.TotalPayment += item.SubTotal
	}

	return data
}
