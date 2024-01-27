package service

import (
	"context"
	"user/internal/model"
)

// UserService defines the interface for the user service
type UserService interface {
	// Create creates a new user and returns its id
	Create(ctx context.Context, user *model.User) (string, error)

	// Get returns a user by id
	Get(ctx context.Context, email string) (*model.User, error)

	// Update updates a user
	Update(ctx context.Context, user *model.UserUpdate) error

	// Delete deletes a user by id
	Delete(ctx context.Context, id string) error
}
