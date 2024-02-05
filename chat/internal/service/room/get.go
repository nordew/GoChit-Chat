package room

import (
	"chat/internal/model"
	chatErrors "chat/pkg/chat_errors"
	"context"
	"errors"
)

func (s *roomService) GetRoomByID(ctx context.Context, id string) (*model.Room, error) {
	const op = "roomService.GetRoomByID"

	room, _, err := s.roomRepo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("failed to get room", op, err.Error())
		if errors.Is(err, chatErrors.ErrRoomNotFound) {
			return nil, err
		}

		return nil, chatErrors.ErrInternal
	}

	return room, nil
}
