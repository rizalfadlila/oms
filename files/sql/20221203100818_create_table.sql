-- +goose Up
CREATE TABLE customers (
                           id bigint NOT NULL PRIMARY KEY,
                           company_name varchar(50) DEFAULT '' NOT NULL,
                           first_name varchar(30) DEFAULT '' NOT NULL,
                           last_name varchar(50) DEFAULT '' NOT NULL,
                           billing_address varchar(225) DEFAULT '' NOT NULL,
                           city varchar(50) DEFAULT '' NOT NULL,
                           state_or_province varchar(20) DEFAULT '' NOT NULL,
                           zip_code varchar(20) DEFAULT '' NOT NULL,
                           email varchar(75) DEFAULT '' NOT NULL,
                           phone_number varchar(30) DEFAULT '' NOT NULL,
                           fax_number varchar(30) DEFAULT '' NOT NULL,
                           ship_address varchar(225) DEFAULT '' NOT NULL,
                           ship_city varchar(50) DEFAULT '' NOT NULL,
                           ship_state_or_province varchar(50) DEFAULT '' NOT NULL,
                           ship_zip_code varchar(20) DEFAULT '' NOT NULL,
                           ship_phone_number varchar(20) DEFAULT '' NOT NULL,
                           created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
                           updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
                           deleted_at timestamp,
                           is_deleted integer DEFAULT 0 NOT NULL
);
CREATE TABLE employees (
                           id bigint NOT NULL PRIMARY KEY,
                           first_name varchar(50) DEFAULT '' NOT NULL,
                           last_name varchar(50) DEFAULT '' NOT NULL,
                           title varchar(50) DEFAULT '' NOT NULL,
                           work_phone varchar(50) DEFAULT '' NOT NULL,
                           created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
                           updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
                           deleted_at timestamp,
                           is_deleted integer DEFAULT 0 NOT NULL
);
CREATE TABLE shipping_methods (
                                  id bigint NOT NULL PRIMARY KEY,
                                  shipping_method varchar(20) DEFAULT '' NOT NULL,
                                  created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
                                  updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
                                  deleted_at timestamp,
                                  is_deleted integer DEFAULT 0 NOT NULL
);
CREATE TABLE products (
                          id bigint NOT NULL PRIMARY KEY,
                          product_name varchar(50) DEFAULT '' NOT NULL,
                          unit_price numeric(20, 2) DEFAULT 0 NOT NULL,
                          in_stock smallint DEFAULT 0 NOT NULL,
                          created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
                          updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
                          deleted_at timestamp,
                          is_deleted integer DEFAULT 0 NOT NULL
);
CREATE TABLE orders (
                        id bigint NOT NULL PRIMARY KEY,
                        customer_id bigint DEFAULT 0 NOT NULL,
                        employee_id bigint DEFAULT 0 NOT NULL,
                        purchase_order_number varchar(30) DEFAULT '' NOT NULL,
                        order_date date DEFAULT CURRENT_TIMESTAMP NOT NULL,
                        ship_date date DEFAULT CURRENT_DATE NOT NULL,
                        shipping_method_id bigint NOT NULL,
                        freight_charge numeric(20, 2) DEFAULT 0 NOT NULL,
                        taxes numeric(20, 2) DEFAULT 0 NOT NULL,
                        payment_received smallint DEFAULT 0 NOT NULL,
                        comment varchar(150) DEFAULT '' NOT NULL,
                        created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
                        updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
                        deleted_at timestamp,
                        is_deleted integer DEFAULT 0 NOT NULL,

                        foreign key (customer_id) references customers(id),
                        foreign key (employee_id) references employees(id),
                        foreign key (shipping_method_id) references shipping_methods(id)
);
CREATE TABLE order_details (
                               id bigint NOT NULL PRIMARY KEY,
                               order_id bigint DEFAULT 0 NOT NULL,
                               product_id bigint DEFAULT 0 NOT NULL,
                               quantity int DEFAULT 0 NOT NULL,
                               unit_price numeric(20, 2) DEFAULT 0 NOT NULL,
                               discount numeric(20, 2) DEFAULT 0 NOT NULL,
                               created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
                               updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
                               deleted_at timestamp,
                               is_deleted integer DEFAULT 0 NOT NULL,

                               foreign key (order_id) references orders(id),
                               foreign key (product_id) references products(id)
);