package customer

const (
	queryStore = `insert into customers (id, company_name, first_name, last_name, billing_address, city, state_or_province, zip_code,
                       email, phone_number, fax_number, ship_address, ship_city, ship_state_or_province, ship_zip_code,
                       ship_phone_number) values (:id, :company_name, :first_name, :last_name, :billing_address, :city, :state_or_province, :zip_code,
                       :email, :phone_number, :fax_number, :ship_address, :ship_city, :ship_state_or_province, :ship_zip_code,
                       :ship_phone_number)`
)

const (
	queryGetIDByEmail = `select id from customers where email = ? and is_deleted = 0`
)
