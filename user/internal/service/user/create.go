package user

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"time"
	"user/internal/model"
	"user/internal/service"
	"user/pkg/auth"
	userErrors "user/pkg/user_errors"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u *userService) Create(ctx context.Context, user *model.User) (*service.CreateUserResponse, *userErrors.CustomErr) {
	const op = "userService.Create"

	err := validateUser(user.Name, user.Email, user.Password)
	if err != nil {
		u.log.Error("error validating user", zap.Error(err), zap.String("op", op))
		return nil, userErrors.New(err, "invalid validation", codes.InvalidArgument)
	}

	hashedPassword, err := u.hasher.Hash(user.Password)
	if err != nil {
		u.log.Error("error hashing password", zap.Error(err), zap.String("op", op))
		return nil, userErrors.NewInternalErr(err)
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
			return nil, userErrors.New(err, err.Error(), codes.AlreadyExists)
		}

		u.log.Error("error creating user", zap.Error(err), zap.String("op", op))
		return nil, userErrors.NewInternalErr(err)
	}

	accessToken, refreshToken, err := u.auth.GenerateTokens(&auth.GenerateTokenClaimsOptions{
		UserId: id.String(),
		Name:   user.Name,
	})
	if err != nil {
		u.log.Error("error generating token", zap.Error(err), zap.String("op", op))
		return nil, userErrors.NewInternalErr(err)
	}

	err = u.userRepo.UpdateRefreshToken(ctx, id.String(), refreshToken)
	if err != nil {
		u.log.Error("error updating refresh token", zap.Error(err), zap.String("op", op))
		return nil, userErrors.NewInternalErr(err)
	}

	resp := &service.CreateUserResponse{
		Id:           id.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return resp, nil
}
