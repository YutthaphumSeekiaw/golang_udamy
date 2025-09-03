-- name: CreateUser :one
INSERT INTO
    users (name, email)
VALUES ($1, $2) RETURNING id,
    name,
    email;

-- name: GetUser :one
SELECT id, name, email FROM users WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT id, name, email FROM users ORDER BY name;