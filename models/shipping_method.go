package models

type ShippingMethod struct {
	ID             int64  `db:"id" json:"id"`
	ShippingMethod string `db:"shipping_method" json:"shipping_method"`
}

func NewShippingMethodFromRowCSV(data interface{}) *ShippingMethod {
	shipping := ShippingMethod{}

	return &shipping
}
