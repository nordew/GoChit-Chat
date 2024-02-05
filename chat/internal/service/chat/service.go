package chat

import (
	"chat/internal/model"
	"chat/internal/repository"
	"chat/internal/service"
	"chat/pkg/logger"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type chatService struct {
	roomRepo repository.RoomRepository
	logger   logger.Logger

	Broadcast      chan *model.Message
	ActiveSessions map[string]*model.User
}

func NewChatService(roomRepo repository.RoomRepository, logger logger.Logger) service.ChatService {
	return &chatService{
		roomRepo: roomRepo,
		logger:   logger,

		Broadcast:      make(chan *model.Message),
		ActiveSessions: make(map[string]*model.User),
	}
}

func (s *chatService) Run() {
	for {
		select {
		case msg := <-s.Broadcast:
			s.handleBroadcast(msg)
		}
	}
}

func (s *chatService) WriteMessage(user *model.User) {
	defer user.Conn.Close()

	for {
		message, ok := <-user.Message
		if !ok {
			return
		}

		user.Conn.WriteJSON(message)
	}
}

func (s *chatService) ReadMessage(user *model.User, roomID string) {
	defer user.Conn.Close()

	for {
		_, m, err := user.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &model.Message{
			Content:   string(m),
			RoomID:    roomID,
			UserID:    user.ID,
			Timestamp: time.Now(),
		}

		s.Broadcast <- msg
	}
}
