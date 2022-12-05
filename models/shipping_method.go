package models

import (
	"github.com/jatis/oms/lib/util"
)

type ShippingMethod struct {
	ID             int64  `db:"id" json:"id"`
	ShippingMethod string `db:"shipping_method" json:"shipping_method"`
}

func NewShippingMethodFromRowCSV(data []interface{}) *ShippingMethod {
	shipping := ShippingMethod{
		ShippingMethod: util.InterfaceToString(data[0]),
	}
	return &shipping
}
