package datatype

import (
	"database/sql/driver"
	"github.com/jatis/oms/lib/util"

	"reflect"
	"time"
)

type TimeInJakarta string

func (t *TimeInJakarta) Value() (driver.Value, error) {
	return t, nil
}

func (t *TimeInJakarta) Scan(value interface{}) error {
	format := "2006-01-02 15:04:05"

	if value == nil {
		*t = ""
		return nil
	}

	if reflect.ValueOf(value).Kind() == reflect.String {
		tm, err := time.Parse(format, value.(string))

		if err != nil {
			return err
		}

		tf := util.TimeInJakarta(tm).Format(format)
		*t = TimeInJakarta(tf[11:])

		return nil
	}

	tm := value.(time.Time)
	tf := util.TimeInJakarta(tm).Format(format)

	*t = TimeInJakarta(tf[11:])

	return nil
}
