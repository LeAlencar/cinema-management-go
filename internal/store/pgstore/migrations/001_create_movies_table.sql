-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS movies (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    "title" VARCHAR(255) NOT NULL,
    "duration" INTEGER NOT NULL,
    "genre" VARCHAR(100) NOT NULL,
    "release_date" TIMESTAMP NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

---- create above / drop below ----
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
DROP TABLE IF EXISTS movies;
