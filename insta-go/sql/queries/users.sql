-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, $2, $2, $3)
RETURNING *;

-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key=$1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUsersByIds :many
SELECT * FROM users WHERE id=ANY($1::UUID[]);