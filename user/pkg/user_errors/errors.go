package userErrors

import "errors"

var (
	ErrEmailAlreadyExists   = errors.New("email already exists")
	ErrInvalidName          = errors.New("invalid name")
	ErrInvalidEmail         = errors.New("invalid email")
	ErrInvalidPassword      = errors.New("invalid password")
	ErrWrongEmailOrPassword = errors.New("wrong email or password")
)
