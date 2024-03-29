-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, user_id, post_text)
VALUES ($1, $2, $2, $3, $4)
RETURNING *;

-- name: GetPostsForUser :many
SELECT * FROM posts
WHERE user_id IN (
	SELECT user_following_id FROM follows
    WHERE follows.user_id = $1
)
ORDER BY created_at DESC;

-- name: GetPostsFromUsers :many
SELECT * FROM posts
WHERE user_id=ANY($1::UUID[])
ORDER BY created_at DESC;
