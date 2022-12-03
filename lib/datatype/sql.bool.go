package datatype

import (
	"database/sql/driver"
	"reflect"
)

type SqlBool bool

func (d *SqlBool) Value() (driver.Value, error) {
	return d, nil
}

func (d *SqlBool) Scan(value interface{}) error {

	if reflect.ValueOf(value).IsZero() {
		*d = false
		return nil
	}

	if reflect.ValueOf(value).Kind() == reflect.Int64 {
		if reflect.ValueOf(value).Int() == 0 {
			*d = false
		} else {
			*d = true
		}
		return nil
	}

	*d = false

	return nil
}
