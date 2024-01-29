package service

import (
	"context"
	"user/internal/model"
)

// UserService defines the interface for the user service
type UserService interface {
	// Create creates a new user and returns its id
	Create(ctx context.Context, user *model.User) (*CreateUserResponse, error)

	// Get returns a user by id
	Get(ctx context.Context, email string) (*model.User, error)

	// Update updates a user
	Update(ctx context.Context, req *UpdateUserRequest) error

	// Delete deletes a user by id
	Delete(ctx context.Context, id string) error

	// Login logs in a user
	Login(ctx context.Context, email string, password string) (string, string, error)
}

type CreateUserResponse struct {
	Id           string
	AccessToken  string
	RefreshToken string
}

type UpdateUserRequest struct {
	User *model.UserUpdate
}
