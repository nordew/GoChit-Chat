package repository

import (
	"context"
	"user/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Get(ctx context.Context, email string) (*model.User, error)
	ChangePassword(ctx context.Context, email, password string) error
}
