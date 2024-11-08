-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS rooms (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    "number" INTEGER NOT NULL UNIQUE,
    "capacity" INTEGER NOT NULL,
    "is_vip" BOOLEAN DEFAULT false,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
DROP TABLE IF EXISTS rooms;