-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
                      user_id UUID UNIQUE,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user IF EXISTS;
-- +goose StatementEnd
