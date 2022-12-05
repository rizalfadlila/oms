package models

type Employee struct {
	ID        int64  `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Title     string `db:"title" json:"title"`
	WorkPhone string `db:"work_phone" json:"work_phone"`
}

func NewEmployeeFromRowCSV(data interface{}) *Employee {
	employee := Employee{}

	return &employee
}
