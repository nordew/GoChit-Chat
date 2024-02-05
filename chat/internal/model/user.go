package model

import "github.com/gorilla/websocket"

type User struct {
	Conn    *websocket.Conn
	Message chan *Message

	ID          string `json:"user_id"`
	Username    string `json:"username"`
	CurrentRoom string `json:"current_room"`
	Rooms       []Room `json:"rooms"`
}
