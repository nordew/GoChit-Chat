-- +goose Up
-- +goose StatementBegin
CREATE TABLE rooms (
                      room_id UUID UNIQUE,
                      room_name VARCHAR(28) UNIQUE NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE room IF EXISTS;
-- +goose StatementEnd
