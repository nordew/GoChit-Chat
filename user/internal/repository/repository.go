package repository

import (
	"context"
	"user/internal/model"
)

type GetFilter struct {
	ID    string
	Email string
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Get(ctx context.Context, filter *GetFilter) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
}
