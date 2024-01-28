package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
	"user/internal/model"
	userErrors "user/pkg/user_errors"
)

func (u *userService) Create(ctx context.Context, user *model.User) (string, error) {
	const op = "userService.Create"

	err := validateUser(user.Name, user.Email, user.Password)
	if err != nil {
		u.log.Error("error validating user", zap.Error(err), zap.String("op", op))
		return "", err
	}

	hashedPassword, err := u.hasher.Hash(user.Password)
	if err != nil {
		u.log.Error("error hashing password", zap.Error(err), zap.String("op", op))
		return "", err
	}

	now := time.Now()
	id := uuid.New()

	parsedUser := &model.User{
		ID:        id.String(),
		Name:      user.Name,
		Email:     user.Email,
		Password:  hashedPassword,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = u.userRepo.Create(ctx, parsedUser)
	if err != nil {
		if errors.Is(err, userErrors.ErrEmailAlreadyExists) {
			return "", userErrors.ErrEmailAlreadyExists
		}

		u.log.Error("error creating user", zap.Error(err), zap.String("op", op))
		return "", err
	}

	return id.String(), nil
}
