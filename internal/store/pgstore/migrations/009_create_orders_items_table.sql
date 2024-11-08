-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS order_items (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    "order_id" uuid NOT NULL REFERENCES orders(id),
    "product_id" uuid NOT NULL REFERENCES products(id),
    "quantity" INTEGER NOT NULL,
    "price" DECIMAL(10,2) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_order_product UNIQUE (order_id, product_id)
);
---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
DROP TABLE IF EXISTS order_items;
