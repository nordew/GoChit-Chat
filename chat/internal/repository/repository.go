package repository

import (
	"chat/internal/model"
	"context"
)

type RoomRepository interface {
	Create(ctx context.Context, room *model.Room) error
	GetByID(ctx context.Context, id string) (*model.Room, bool, error)
	GetAllUsersFromRoom(ctx context.Context, roomID string) ([]model.User, error)
	Update(ctx context.Context, id string, room *model.Room) error
	AddUser(ctx context.Context, roomID, userID string) error
	DeleteUser(ctx context.Context, roomID, userID string) error
}
