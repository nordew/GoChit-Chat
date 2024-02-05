package room

import (
	"chat/internal/service"
	chatErrors "chat/pkg/chat_errors"
	"context"
)

func (s *roomService) UpdateRoom(ctx context.Context, updateRoomReq *service.UpdateRoomReq) error {
	const op = "roomService.UpdateRoom"

	if err := s.roomRepo.Update(ctx, updateRoomReq.ID, updateRoomReq.Room); err != nil {
		s.logger.Error("failed to update room", op, err.Error())
		return chatErrors.ErrInternal
	}

	return nil
}

func (s *roomService) AddUser(ctx context.Context, roomID, userID string) error {
	const op = "roomService.AddUser"

	if err := s.roomRepo.AddUser(ctx, roomID, userID); err != nil {
		s.logger.Error("failed to add user to room", op, err.Error())
		return chatErrors.ErrInternal
	}

	return nil
}
