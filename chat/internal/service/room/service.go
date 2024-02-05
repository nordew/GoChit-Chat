package room

import (
	"chat/internal/repository"
	"chat/internal/service"
	"chat/pkg/logger"
)

type roomService struct {
	roomRepo repository.RoomRepository
	logger   logger.Logger
}

func NewRoomService(roomRepo repository.RoomRepository, logger logger.Logger) service.RoomService {
	return &roomService{
		roomRepo: roomRepo,
		logger:   logger,
	}
}
