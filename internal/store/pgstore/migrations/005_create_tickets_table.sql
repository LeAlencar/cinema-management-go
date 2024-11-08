-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS tickets (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    "session_id" uuid NOT NULL REFERENCES sessions(id),
    "customer_id" uuid NOT NULL REFERENCES customers(id),
    "seat_number" VARCHAR(10) NOT NULL,
    "price" DECIMAL(10,2) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_seat_session UNIQUE (session_id, seat_number)
);

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
DROP TABLE IF EXISTS tickets;