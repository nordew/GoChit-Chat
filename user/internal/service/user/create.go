package user

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
	"user/internal/model"
)

func (u *userService) Create(ctx context.Context, user *model.User) (string, error) {
	const op = "userService.Create"

	err := validateUser(user.Name, user.Email, user.Password)
	if err != nil {
		u.log.Error("error validating user", zap.Error(err), zap.String("op", op))
		return "", err
	}

	exists, err := u.userRepo.CheckIfEmailExists(ctx, user.Email)
	if err != nil {
		u.log.Error("error checking if email exists", zap.Error(err), zap.String("op", op))
		return "", err
	}

	if exists {
		u.log.Error("email already exists", zap.String("op", op))
		return "", ErrEmailAlreadyExists
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
		u.log.Error("error creating user", zap.Error(err), zap.String("op", op))
		return "", err
	}

	return id.String(), nil
}
