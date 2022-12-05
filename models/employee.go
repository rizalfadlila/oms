package models

import "github.com/jatis/oms/lib/util"

type Employee struct {
	ID        int64  `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Title     string `db:"title" json:"title"`
	WorkPhone string `db:"work_phone" json:"work_phone"`
}

func NewEmployeeFromRowCSV(data []interface{}) *Employee {
	employee := Employee{
		FirstName: util.InterfaceToString(data[0]),
		LastName:  util.InterfaceToString(data[1]),
		Title:     util.InterfaceToString(data[2]),
		WorkPhone: util.InterfaceToString(data[3]),
	}

	return &employee
}
