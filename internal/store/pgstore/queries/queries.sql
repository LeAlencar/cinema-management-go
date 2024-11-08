-- name: GetMovie :one
SELECT * FROM movies WHERE id = $1;

-- name: ListMovies :many
SELECT * FROM movies;

-- name: ListMoviesByGenre :many
SELECT * FROM movies WHERE genre = $1;

-- name: CreateMovie :one
INSERT INTO movies (
    title,
    duration,
    genre,
    release_date
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateMovie :one
UPDATE movies
SET
    title = COALESCE($2, title),
    duration = COALESCE($3, duration),
    genre = COALESCE($4, genre),
    release_date = COALESCE($5, release_date),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteMovie :exec
DELETE FROM movies WHERE id = $1;

