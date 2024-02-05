package chatErrors

import "errors"

var (
	ErrInternal     = errors.New("server internal error")
	ErrRoomNotFound = errors.New("room doesn't exist")
)
