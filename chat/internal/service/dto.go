package service

import "chat/internal/model"

type CreateRoomReq struct {
	RoomName string
	Username string
	UserID   string
}

type UpdateRoomReq struct {
	ID   string
	Room *model.Room
}
