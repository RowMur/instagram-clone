-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, user_id, post_text)
VALUES ($1, $2, $2, $3, $4)
RETURNING *;

-- name: GetPostsFromUsers :many
SELECT * FROM posts
WHERE user_id=ANY($1::UUID[])
ORDER BY created_at DESC;