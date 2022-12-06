package models

type ReportOrder struct {
	CustomerName   string `db:"customer_name"`
	EmployeeName   string `db:"employee_name"`
	ShippingMethod string `db:"shipping_method"`
}

type ReportOrderDetail struct {
	ID          int64   `db:"product_id" json:"product_id"`
	ProductName string  `db:"product_name" json:"product_name"`
	Quantity    int     `db:"quantity" json:"quantity"`
	UnitPrice   float64 `db:"unit_price" json:"unit_price"`
	Discount    float64 `db:"discount" json:"discount"`
	SubTotal    float64 `db:"-" json:"sub_total"`
}

type ResponseReportOrder struct {
	CustomerName   string              `json:"customer_name"`
	EmployeeName   string              `json:"employee_name"`
	ShippingMethod string              `json:"shipping_method_name"`
	Items          []ReportOrderDetail `json:"items"`
	TotalPayment   float64             `json:"total_payment"`
}

func ComposeReportData(order ReportOrder, detail []ReportOrderDetail) *ResponseReportOrder {
	data := &ResponseReportOrder{
		CustomerName:   order.CustomerName,
		EmployeeName:   order.EmployeeName,
		ShippingMethod: order.ShippingMethod,
	}

	items := make([]ReportOrderDetail, 0)
	for _, item := range detail {
		item.SubTotal = (item.UnitPrice * float64(item.Quantity)) - item.Discount

		items = append(items, item)

		data.TotalPayment += item.SubTotal
	}

	data.Items = items

	return data
}
