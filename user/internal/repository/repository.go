package repository

import (
	"context"
	"user/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Get(ctx context.Context, id string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}
