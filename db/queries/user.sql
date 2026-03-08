-- name: GetUserByID :one
SELECT id, email, name, password, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, email, name, password, created_at, updated_at
FROM users
WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users (email, name, password)
VALUES ($1, $2, $3)
RETURNING id, email, name, password, created_at, updated_at;

-- name: ListUsers :many
SELECT id, email, name, password, created_at, updated_at
FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
