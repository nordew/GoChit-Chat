package v1

import (
	"chat/internal/service"
	"chat/pkg/logger"

	"github.com/gin-gonic/gin"
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
)

type Handler struct {
	roomService service.RoomService
	chatService service.ChatService
	userClient  desc.UserClient
	logger      logger.Logger
}

func NewHandler(roomService service.RoomService, userClient desc.UserClient, chatService service.ChatService, logger logger.Logger) *Handler {
	return &Handler{
		roomService: roomService,
		userClient:  userClient,
		chatService: chatService,
		logger:      logger,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	router.POST("/create-room", h.createRoom)
	router.GET("/ws/join-room/:roomID", h.joinRoom)

	return router
}
