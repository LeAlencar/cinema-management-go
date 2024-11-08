-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS orders (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    "customer_id" uuid NOT NULL REFERENCES customers(id),
    "total_amount" DECIMAL(10,2) NOT NULL,
    "status" VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
DROP TABLE IF EXISTS orders;
