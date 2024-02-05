package room

import (
	"context"
)

func (r roomRepository) DeleteUser(ctx context.Context, roomID, userID string) error {
	query := "DELETE FROM user_to_room WHERE user_id = $1 AND room_id = $2"
	_, err := r.conn.Exec(ctx, query, userID, roomID)
	if err != nil {
		return err
	}

	return nil
}
