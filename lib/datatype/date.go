package datatype

import (
	"database/sql/driver"
	"github.com/jatis/oms/lib/util"
	"reflect"
	"time"
)

type DateInJakarta string

func (d *DateInJakarta) Value() (driver.Value, error) {
	return d, nil
}

func (d *DateInJakarta) Scan(value interface{}) error {
	format := "2006-01-02 15:04:05"

	if value == nil {
		*d = ""
		return nil
	}

	if reflect.ValueOf(value).Kind() == reflect.String {
		date, err := time.Parse(format, value.(string))

		if err != nil {
			return err
		}

		dt := util.TimeInJakarta(date).Format(format)
		*d = DateInJakarta(dt[0:10])

		return nil
	}

	date := value.(time.Time)
	dt := util.TimeInJakarta(date).Format(format)
	*d = DateInJakarta(dt[0:10])

	return nil
}

func (d *DateInJakarta) String() string {
	if d == nil {
		return ""
	}
	return string(*d)
}
