-- +goose Up
CREATE TABLE follows (
    user_id UUID REFERENCES users(id) NOT NULL,
    user_following_id UUID REFERENCES users(id) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    PRIMARY KEY(user_id, user_following_id)
);

-- +goose Down
DROP TABLE follows;