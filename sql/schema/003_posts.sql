-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID REFERENCES users(id) NOT NULL,
    post_text TEXT NOT NULL
);

-- +goose Down
DROP TABLE posts;