package user

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"time"
	"user/internal/model"
	"user/pkg/auth"
	userErrors "user/pkg/user_errors"

	"go.uber.org/zap"
)

func (u userService) Get(ctx context.Context, email string) (*model.User, *userErrors.CustomErr) {
	const op = "userService.Get"

	user, err := u.userRepo.GetByEmail(ctx, email)

	if err != nil {
		u.log.Error("error getting user", zap.Error(err), zap.String("op", op))
		return nil, userErrors.NewInternalErr(err)
	}

	return user, nil
}

func (u userService) Login(ctx context.Context, email string, password string) (string, string, *userErrors.CustomErr) {
	const op = "userService.Login"

	valid := IsValidEmail(email)
	if !valid {
		return "", "", userErrors.New(userErrors.ErrInvalidEmail, "invalid email", codes.InvalidArgument)
	}

	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, userErrors.ErrWrongEmailOrPassword) {
			return "", "", userErrors.PasswordOrEmailMismatch(err)
		}

		u.log.Error("error getting user", zap.Error(err), zap.String("op", op))
		return "", "", userErrors.NewInternalErr(err)
	}

	err = u.hasher.Compare(user.Password, password)
	if err != nil {
		u.log.Error("error comparing password", zap.Error(err), zap.String("op", op))
		return "", "", userErrors.PasswordOrEmailMismatch(err)
	}

	accessToken, refreshToken, err := u.auth.GenerateTokens(&auth.GenerateTokenClaimsOptions{
		UserId: user.ID,
	})
	if err != nil {
		u.log.Error("error generating token", zap.Error(err), zap.String("op", op))
		return "", "", userErrors.NewInternalErr(err)
	}

	var userWithUpdatedTokens = &model.User{
		ID:           user.ID,
		RefreshToken: refreshToken,
		UpdatedAt:    time.Now(),
	}

	err = u.userRepo.Update(ctx, userWithUpdatedTokens)
	if err != nil {
		u.log.Error("error updating user", zap.Error(err), zap.String("op", op))
		return "", "", userErrors.NewInternalErr(err)
	}

	return accessToken, refreshToken, nil
}
