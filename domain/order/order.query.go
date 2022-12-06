package order

const (
	queryStoreOrder = `insert into orders (id, customer_id, employee_id, purchase_order_number, order_date, ship_date, shipping_method_id,
                    freight_charge, taxes, payment_received, comment) values (:id, :customer_id, :employee_id, :purchase_order_number, :order_date, :ship_date, :shipping_method_id,
                    :freight_charge, :taxes, :payment_received, :comment)`

	queryStoreOrderDetail = `insert into order_details (id, order_id, product_id, quantity, unit_price, discount) values (:id, :order_id, :product_id, :quantity, :unit_price, :discount)`
)

const (
	queryGetReportOrder = `select CONCAT(c.first_name, c.last_name) as customer_name, CONCAT(e.first_name, e.last_name) as employee_name, s.shipping_method
									from orders o
									inner join customers c on c.id = o.customer_id and c.is_deleted = 0
									inner join employee e on e.id = o.employee_id and e.is_deleted = 0
									inner join shipping_methods s on s.id = o.shipping_method_id and s.is_deleted = 0
									where od.order_id = ? and o.is_deleted = 0`

	queryGetReportOrderDetail = `select p.id as product_id, p.product_name, od.quantity, od.unit_price, od.discount 
									from order_detail od
									inner join products p on p.id = od.product_id and o.is_deleted = 0
									where od.order_id = ? and od.is_deleted = 0`

	queryGetIDByPO = `select id from orders where purchase_order_number = ? and is_deleted = 0`
)
