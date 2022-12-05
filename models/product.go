package models

import (
	"github.com/jatis/oms/lib/util"
)

type Product struct {
	ID          int64   `db:"id" json:"id"`
	ProductName string  `db:"product_name" json:"product_name"`
	UnitPrice   float64 `db:"unit_price" json:"unit_price"`
	InStock     int16   `db:"in_stock" json:"in_stock"`
}

func NewProductFromRowCSV(data []interface{}) *Product {
	product := Product{
		ProductName: util.InterfaceToString(data[0]),
		UnitPrice:   util.InterfaceToFloat64(data[1]),
		InStock:     int16(util.InterfaceToInt(data[2])),
	}

	return &product
}
