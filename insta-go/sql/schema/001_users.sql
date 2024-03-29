-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    api_key VARCHAR(64) NOT NULL UNIQUE DEFAULT (
        encode(sha256(random()::text::bytea), 'hex')
    )
);

-- +goose Down
DROP TABLE users;