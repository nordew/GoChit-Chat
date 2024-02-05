package v1

import (
	"chat/internal/model"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"net/http"
)

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) joinRoom(c *gin.Context) {
	roomID := c.Param("roomID")

	token := c.GetHeader("Authorization")

	resp, err := h.userClient.ParseAccessToken(context.Background(), &desc.ParseAccessTokenRequest{
		Token: token,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, "failed to parse token")
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	user := &model.User{
		Conn:        conn,
		Message:     make(chan *model.Message, 10),
		ID:          resp.UserID,
		Username:    resp.UserID,
		CurrentRoom: roomID,
	}

	message := model.JoinRoomMessage(user.ID, user.CurrentRoom)

	if err := h.chatService.RegisterUser(context.Background(), user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		conn.WriteJSON(err.Error())
		conn.Close()
		return
	}

	h.chatService.BroadcastMessage(context.Background(), message)

	go h.chatService.WriteMessage(user)
	h.chatService.ReadMessage(user, roomID)
}
