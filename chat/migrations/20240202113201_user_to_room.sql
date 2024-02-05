-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_to_room (
                              id SERIAL PRIMARY KEY,
                              user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
                              room_id UUID REFERENCES rooms(room_id) ON DELETE CASCADE,

                              UNIQUE (user_id, room_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_to_room IF EXISTS;
-- +goose StatementEnd
