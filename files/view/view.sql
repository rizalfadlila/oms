# List of customers located in Irvine city
SELECT
    id,
    company_name,
    first_name,
    last_name,
    billing_address,
    city,
    state_or_province,
    zip_code,
    email,
    phone_number,
    fax_number,
    ship_address,
    ship_city,
    ship_state_or_province,
    ship_zip_code,
    ship_phone_number
FROM customers
    WHERE
            city = 'Irvine city'
        AND is_deleted = 0;

# List of customers whose order is handled by an employee named Adam Barr
SELECT
    c.id,
    c.company_name,
    c.first_name,
    c.last_name,
    c.billing_address,
    c.city,
    c.state_or_province,
    c.zip_code,
    c.email,
    c.phone_number,
    c.fax_number,
    c.ship_address,
    c.ship_city,
    c.ship_state_or_province,
    c.ship_zip_code,
    c.ship_phone_number
FROM customers c
INNER JOIN orders o ON o.customer_id = c.id and o.is_deleted = 0
INNER JOIN employees e on e.id = o.employee_id and e.is_deleted = 0
WHERE
            e.first_name = 'Adam'
        AND e.last_name = 'Barr'
        AND c.is_deleted = 0;

# List of products which are ordered by "Contonso, Ltd" Company
SELECT
    p.product_name,
    p.unit_price,
    p.in_stock
FROM products p
INNER JOIN order_details od on p.id = od.product_id and od.is_deleted = 0
INNER JOIN orders o on o.id = od.order_id and o.is_deleted = 0
INNER JOIN customers c on o.customer_id = c.id and c.is_deleted = 0
WHERE
            c.company_name = 'Contonso, Ltd'
        AND p.is_deleted = 0;

# List of transactions (orders) which has "UPS Ground" as shipping method.
SELECT
    o.id,
    o.customer_id,
    o.employee_id,
    o.purchase_order_number,
    o.order_date,
    o.ship_date,
    o.shipping_method_id,
    o.freight_charge,
    o.taxes,
    o.payment_received,
    o.comment
FROM orders o
INNER JOIN shipping_methods sm on o.shipping_method_id = sm.id and sm.is_deleted = 0
WHERE
        sm.shipping_method = 'UPS Ground'
    AND o.is_deleted = 0;

# List of total cost (including tax and freight charge) for every order sorted by ship date
SELECT
    od.id, o.ship_date, SUM(((od.unit_price * od.quantity) - od.discount) + (select SUM((o.freight_charge + o.taxes))
                                                                             from
                                                                                 orders o
                                                                             where o.id = od.order_id
                                                                             group by o.id)
        )  total_cost
FROM order_details od
INNER JOIN orders o on od.order_id = o.id and o.is_deleted = 0
GROUP BY o.id, o.ship_date
ORDER BY o.ship_date