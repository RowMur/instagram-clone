-- name: CreateFollow :one
INSERT INTO follows (user_id, user_following_id, created_at)
VALUES ($1, $2, $3)
RETURNING *;