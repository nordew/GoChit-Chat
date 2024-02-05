package room

import (
	"chat/internal/repository"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type roomRepository struct {
	conn *pgxpool.Pool
}

func NewRoomRepository(conn *pgxpool.Pool) repository.RoomRepository {
	return &roomRepository{
		conn: conn,
	}
}

func (r roomRepository) userExists(ctx context.Context, tx pgx.Tx, userID, roomID, op string) (bool, error) {
	var exists bool

	query := "SELECT EXISTS(SELECT 1 FROM user_to_room WHERE user_id = $1 AND room_id = $2)"

	err := tx.QueryRow(ctx, query, userID, roomID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
