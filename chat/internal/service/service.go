package service

import (
	"chat/internal/model"
	"context"
)

type RoomService interface {
	CreateRoom(ctx context.Context, createRoomReq *CreateRoomReq) (string, error)
	UpdateRoom(ctx context.Context, updateRoomReq *UpdateRoomReq) error
	GetRoomByID(ctx context.Context, id string) (*model.Room, error)
	AddUser(ctx context.Context, roomID, userID string) error
}

type ChatService interface {
	Run()
	RegisterUser(ctx context.Context, user *model.User) error
	UnregisterUser(ctx context.Context, user *model.User) error
	BroadcastMessage(ctx context.Context, message *model.Message)
	WriteMessage(user *model.User)
	ReadMessage(user *model.User, roomID string)
}
