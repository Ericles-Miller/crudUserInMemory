-- name: FindAll :many
SELECT id, first_name, last_name, biography
FROM users;

-- name: FindById :one
SELECT id, first_name, last_name, biography
FROM users
WHERE id = $1;

-- name: Insert :one
INSERT INTO users (first_name, last_name, biography)
VALUES ($1, $2, $3)
RETURNING id, first_name, last_name, biography;

-- name: Update :one
UPDATE users
SET first_name = $2,
    last_name  = $3,
    biography  = $4
WHERE id = $1
RETURNING id, first_name, last_name, biography;

-- name: Delete :one
DELETE FROM users
WHERE id = $1
RETURNING id, first_name, last_name, biography;
