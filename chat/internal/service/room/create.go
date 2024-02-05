package room

import (
	"chat/internal/model"
	"chat/internal/service"
	chatErrors "chat/pkg/chat_errors"
	"context"
	"github.com/google/uuid"
)

func (s *roomService) CreateRoom(ctx context.Context, req *service.CreateRoomReq) (string, error) {
	const op = "roomService.CreateRoom"

	id := uuid.New().String()

	users := append([]model.User{}, model.User{Username: req.Username, ID: req.UserID})

	room := &model.Room{
		ID:    id,
		Name:  req.RoomName,
		Users: users,
	}

	if err := s.roomRepo.Create(ctx, room); err != nil {
		s.logger.Error("failed to create room", op, err.Error())
		return "", chatErrors.ErrInternal
	}

	return id, nil
}
