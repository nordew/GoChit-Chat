package v1

import (
	"chat/internal/controller/http/dto"
	"chat/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"net/http"
)

func (h *Handler) createRoom(c *gin.Context) {
	token := c.GetHeader("Authorization")

	var req dto.CreateRoomRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "invalid json")
		return
	}

	resp, err := h.userClient.ParseAccessToken(context.Background(), &desc.ParseAccessTokenRequest{
		Token: token,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.roomService.CreateRoom(context.Background(), &service.CreateRoomReq{
		RoomName: req.Name,
		Username: resp.Name,
		UserID:   resp.UserID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}
