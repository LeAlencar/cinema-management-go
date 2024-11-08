-- Write your migrate up statements here

CREATE INDEX idx_employees_email ON employees("email");
CREATE INDEX idx_employees_role ON employees("role");
CREATE INDEX idx_products_category ON products("category");
CREATE INDEX idx_orders_customer_id ON orders("customer_id");
CREATE INDEX idx_orders_status ON orders("status");
CREATE INDEX idx_order_items_order_id ON order_items("order_id");
CREATE INDEX idx_order_items_product_id ON order_items("product_id");
---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
DROP INDEX IF EXISTS idx_order_items_product_id;
DROP INDEX IF EXISTS idx_order_items_order_id;
DROP INDEX IF EXISTS idx_orders_status;
DROP INDEX IF EXISTS idx_orders_customer_id;
DROP INDEX IF EXISTS idx_products_category;
DROP INDEX IF EXISTS idx_employees_role;
DROP INDEX IF EXISTS idx_employees_email;
