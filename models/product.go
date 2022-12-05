package models

import "github.com/jatis/oms/lib/datatype"

type Product struct {
	ID          int64            `db:"id" json:"id"`
	ProductName string           `db:"product_name" json:"product_name"`
	UnitPrice   float64          `db:"unit_price" json:"unit_price"`
	InStock     datatype.SqlBool `db:"in_stock" json:"in_stock"`
}

func NewProductFromRowCSV(data interface{}) *Product {
	product := Product{}

	return &product
}
