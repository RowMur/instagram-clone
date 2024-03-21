-- name: CreateFollow :one
INSERT INTO follows (user_id, user_following_id, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetFollowsByUser :many
SELECT * FROM follows WHERE user_id=$1;

-- name: GetFollowersByUser :many
SELECT * FROM follows WHERE user_following_id=$1;