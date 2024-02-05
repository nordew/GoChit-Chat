package chat

import (
	"chat/internal/model"
	"context"
	"fmt"
)

func (s *chatService) handleBroadcast(message *model.Message) {
	const op = "chatService.handleBroadcast"

	for _, user := range s.ActiveSessions {
		user.Message <- message
	}
}

func (s *chatService) RegisterUser(ctx context.Context, user *model.User) error {
	const op = "chatService.RegisterUser"

	_, exists, err := s.roomRepo.GetByID(ctx, user.CurrentRoom)
	if err != nil {
		s.logger.Error("failed to get room", op, err.Error())
		return err
	}

	if !exists {
		if err := s.roomRepo.AddUser(ctx, user.CurrentRoom, user.ID); err != nil {
			s.logger.Error("failed to add user", op, err.Error())
			return err
		}
	}

	s.ActiveSessions[user.ID] = user

	return nil
}

func (s *chatService) UnregisterUser(ctx context.Context, user *model.User) error {
	const op = "chatService.UnregisterUSer"

	if err := s.roomRepo.DeleteUser(ctx, user.CurrentRoom, user.ID); err != nil {
		s.logger.Error("failed to delete user", op, err.Error())
		return err
	}

	message := &model.Message{
		UserID:  user.ID,
		RoomID:  user.CurrentRoom,
		Content: fmt.Sprintf("User %s left the room", user.Username),
	}

	s.Broadcast <- message

	return nil
}

func (s *chatService) BroadcastMessage(ctx context.Context, message *model.Message) {
	s.Broadcast <- message
}
