package repository

import (
	"context"
	"user/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error

	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetById(ctx context.Context, id string) (*model.User, error)

	Update(ctx context.Context, user *model.User) error
	UpdateRefreshToken(ctx context.Context, id string, refreshToken string) error

	Delete(ctx context.Context, id string) error
}
