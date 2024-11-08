-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS sessions (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    "movie_id" uuid NOT NULL REFERENCES movies(id),
    "room_id" uuid NOT NULL REFERENCES rooms(id),
    "start_time" TIMESTAMP NOT NULL,
    "price" DECIMAL(10,2) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_room_time UNIQUE (room_id, start_time)
);
---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
DROP TABLE IF EXISTS sessions;