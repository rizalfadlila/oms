package models

import (
	"github.com/jatis/oms/lib/util"
)

type Customers struct {
	ID                  int64  `db:"id" json:"id"`
	CompanyName         string `db:"company_name" json:"company_name"`
	FirstName           string `db:"first_name" json:"first_name"`
	LastName            string `db:"last_name" json:"last_name"`
	BillingAddress      string `db:"billing_address" json:"billing_address"`
	City                string `db:"city" json:"city"`
	StateOrProvince     string `db:"state_or_province" json:"state_or_province"`
	ZipCode             string `db:"zip_code" json:"zip_code"`
	Email               string `db:"email" json:"email"`
	PhoneNumber         string `db:"phone_number" json:"phone_number"`
	FaxNumber           string `db:"fax_number" json:"fax_number"`
	ShipAddress         string `db:"ship_address" json:"ship_address"`
	ShipCity            string `db:"ship_city" json:"ship_city"`
	ShipStateOrProvince string `db:"ship_state_or_province" json:"ship_state_or_province"`
	ShipZipCode         string `db:"ship_zip_code" json:"ship_zip_code"`
	ShipPhoneNumber     string `db:"ship_phone_number" json:"ship_phone_number"`
}

func NewCustomerFromRowCSV(data []interface{}) *Customers {
	customer := Customers{
		CompanyName:         util.InterfaceToString(data[0]),
		FirstName:           util.InterfaceToString(data[1]),
		LastName:            util.InterfaceToString(data[2]),
		BillingAddress:      util.InterfaceToString(data[3]),
		City:                util.InterfaceToString(data[4]),
		StateOrProvince:     util.InterfaceToString(data[5]),
		ZipCode:             util.InterfaceToString(data[6]),
		Email:               util.InterfaceToString(data[7]),
		PhoneNumber:         util.InterfaceToString(data[8]),
		FaxNumber:           util.InterfaceToString(data[9]),
		ShipAddress:         util.InterfaceToString(data[10]),
		ShipCity:            util.InterfaceToString(data[11]),
		ShipStateOrProvince: util.InterfaceToString(data[12]),
		ShipZipCode:         util.InterfaceToString(data[13]),
		ShipPhoneNumber:     util.InterfaceToString(data[14]),
	}

	return &customer
}
