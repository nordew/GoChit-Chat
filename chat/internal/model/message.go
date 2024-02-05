package model

import (
	"fmt"
	"time"
)

type Message struct {
	UserID    string    `json:"user_id"`
	RoomID    string    `json:"room_id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

func JoinRoomMessage(userID, roomID string) *Message {
	return &Message{
		UserID:    userID,
		RoomID:    roomID,
		Content:   fmt.Sprintf("User %s joined the room", userID),
		Timestamp: time.Now(),
	}
}
