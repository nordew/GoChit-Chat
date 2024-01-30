package service

import (
	"context"
	"user/internal/model"
	userErrors "user/pkg/user_errors"
)

// UserService defines the interface for the user service
type UserService interface {
	// Create creates a new user and returns its id
	Create(ctx context.Context, user *model.User) (*CreateUserResponse, *userErrors.CustomErr)

	// Get returns a user by id
	Get(ctx context.Context, email string) (*model.User, *userErrors.CustomErr)

	// Update updates a user
	Update(ctx context.Context, req *UpdateUserRequest) *userErrors.CustomErr

	// Delete deletes a user by id
	Delete(ctx context.Context, id string) *userErrors.CustomErr

	// Login logs in a user
	Login(ctx context.Context, email string, password string) (string, string, *userErrors.CustomErr)
}

type CreateUserResponse struct {
	Id           string
	AccessToken  string
	RefreshToken string
}

type UpdateUserRequest struct {
	User *model.UserUpdate
}
