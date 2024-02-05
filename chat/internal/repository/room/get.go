package room

import (
	"chat/internal/model"
	chatErrors "chat/pkg/chat_errors"
	"context"
	"database/sql"
)

func (r roomRepository) GetByID(ctx context.Context, id string) (*model.Room, bool, error) {
	room := &model.Room{}

	query := "SELECT room_id, room_name FROM rooms WHERE room_id = $1"
	err := r.conn.QueryRow(ctx, query, id).Scan(&room.ID, &room.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, chatErrors.ErrRoomNotFound
		}
		return nil, false, err
	}

	query = "SELECT user_id FROM user_to_room WHERE room_id = $1"
	rows, err := r.conn.Query(ctx, query, id)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()

	for rows.Next() {
		var userID string
		err := rows.Scan(&userID)
		if err != nil {
			return nil, false, err
		}

		user := model.User{ID: userID}
		room.Users = append(room.Users, user)
	}

	return room, true, nil
}

func (r roomRepository) GetAllUsersFromRoom(ctx context.Context, roomID string) ([]model.User, error) {
	const op = "roomRepo.GetAllUsersFromRoom"

	query := `
		SELECT u.user_id, u.username
		FROM users u
		INNER JOIN user_to_room ur ON u.user_id = ur.user_id
		WHERE ur.room_id = $1
	`

	rows, err := r.conn.Query(ctx, query, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
