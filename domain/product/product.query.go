package product

const (
	queryStore = `insert into products (id, product_name, unit_price, in_stock) values (:id, :product_name, :unit_price, :in_stock)`
)
