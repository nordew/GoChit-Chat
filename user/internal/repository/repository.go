package repository

import (
	"context"
	"github.com/google/uuid"
	"user/internal/model"
)

type GetFilter struct {
	ID    uuid.UUID
	Email string
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Get(ctx context.Context, filter *GetFilter) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
}
