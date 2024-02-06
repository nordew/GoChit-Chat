package userErrors

import (
	"errors"
	"google.golang.org/grpc/codes"
)

var (
	ErrEmailAlreadyExists   = errors.New("email already exists")
	ErrInvalidName          = errors.New("invalid name")
	ErrInvalidEmail         = errors.New("invalid email")
	ErrInvalidPassword      = errors.New("invalid password")
	ErrWrongEmailOrPassword = errors.New("wrong email or password")
	ErrUserNotFound         = errors.New("user not found")
)

type CustomErr struct {
	Err  error
	Msg  string
	Code codes.Code
}

func New(err error, msg string, code codes.Code) *CustomErr {
	return &CustomErr{
		Err:  err,
		Msg:  msg,
		Code: code,
	}
}

func NewInternalErr(err error) *CustomErr {
	return &CustomErr{
		Err:  err,
		Msg:  "internal sever error",
		Code: codes.Internal,
	}
}

func PasswordOrEmailMismatch(err error) *CustomErr {
	return &CustomErr{
		Err:  err,
		Msg:  ErrWrongEmailOrPassword.Error(),
		Code: codes.NotFound,
	}
}
