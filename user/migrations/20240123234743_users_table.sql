-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_table (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    refresh_token VARCHAR NOT NULL DEFAULT '',
    role VARCHAR(50) NOT NULL CHECK (role IN ('USER', 'ADMIN')) DEFAULT 'USER',
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_table IF EXISTS;
-- +goose StatementEnd