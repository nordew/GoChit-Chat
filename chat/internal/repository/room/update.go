package room

import (
	"chat/internal/model"
	"context"
)

func (r roomRepository) Update(ctx context.Context, id string, room *model.Room) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := "UPDATE rooms SET room_name = $2 WHERE room_id = $1"
	_, err = tx.Exec(ctx, query, id, room.Name)
	if err != nil {
		return err
	}

	query = "DELETE FROM user_to_room WHERE room_id = $1"
	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	for _, client := range room.Users {
		query = "INSERT INTO user_to_room (user_id, room_id) VALUES ($1, $2)"
		_, err = tx.Exec(ctx, query, client.ID, id)
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

func (r roomRepository) AddUser(ctx context.Context, roomID, userID string) error {
	query := "INSERT INTO user_to_room (room_id, user_id) VALUES ($1, $2)"

	_, err := r.conn.Exec(ctx, query, roomID, userID)
	if err != nil {
		return err
	}

	return nil
}
