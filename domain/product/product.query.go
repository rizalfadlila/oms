package product

const (
	queryStore = `insert into products (id, product_name, unit_price, in_stock) values (:id, :product_name, :unit_price, :in_stock)`
)

const (
	queryGetByProductName = `select id, unit_price from products where product_name = ? and is_deleted = 0 and in_stock = 1`
)
