package shippingmethod

const (
	queryStore = `insert into shipping_methods (id, shipping_method) values (:id, :shipping_method)`
)

const (
	queryGetIDByMethod = `select id from shipping_methods where shipping_method = ? and is_deleted = 0`
)
