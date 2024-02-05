package room

import (
	"chat/internal/model"
	"context"
)

func (r roomRepository) Create(ctx context.Context, room *model.Room) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := "INSERT INTO rooms (room_id, room_name) VALUES ($1, $2)"
	_, err = tx.Exec(ctx, query, room.ID, room.Name)
	if err != nil {
		return err
	}

	for _, user := range room.Users {
		query := "INSERT INTO users (user_id, username) VALUES ($1, $2)"
		_, err := tx.Exec(ctx, query, user.ID, user.Username)
		if err != nil {
			return err
		}
	}

	for _, user := range room.Users {
		query = "INSERT INTO user_to_room (user_id, room_id) VALUES ($1, $2)"
		_, err = tx.Exec(ctx, query, user.ID, room.ID)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
