package user

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"time"
	"unicode/utf8"
	"user/internal/model"
	"user/internal/service"
	"user/pkg/auth"
	userErrors "user/pkg/user_errors"
)

func (u *userService) Refresh(ctx context.Context, token string) (string, string, *userErrors.CustomErr) {
	const op = "userService.Refresh"

	user, err := u.userRepo.GetByToken(ctx, token)
	if err != nil {
		if errors.Is(err, userErrors.ErrUserNotFound) {
			u.log.Error("user not found", zap.String("op", op))
			return "", "", userErrors.New(err, "invalid refresh token", codes.InvalidArgument)
		}
	}

	accessToken, refreshToken, err := u.auth.GenerateTokens(&auth.GenerateTokenClaimsOptions{
		UserId: user.ID,
		Name:   user.Name,
	})
	if err != nil {
		u.log.Error("failed to generate tokens", zap.Error(err), zap.String("op", op))
		return "", "", userErrors.NewInternalErr(err)
	}

	if err := u.userRepo.UpdateRefreshToken(ctx, user.ID, refreshToken); err != nil {
		u.log.Error("failed to update refresh token", zap.Error(err), zap.String("op", op))
		return "", "", userErrors.NewInternalErr(err)
	}

	return accessToken, refreshToken, nil
}

func (u *userService) Update(ctx context.Context, request *service.UpdateUserRequest) *userErrors.CustomErr {
	const op = "userService.Update"

	oldUser, err := u.userRepo.GetById(ctx, request.User.ID)
	if err != nil {
		u.log.Error("error getting user", zap.Error(err), zap.String("op", op))
		return userErrors.NewInternalErr(err)
	}

	if request.User.Password {
		err := u.updateUserPassword(ctx, oldUser, request.User.OldPassword, request.User.NewPassword)
		if err != nil {
			return err
		}
	}

	if request.User.Name != "" {
		err := u.updateUserName(ctx, oldUser, request.User.Name)
		if err != nil {
			return err
		}
	}

	if request.User.Email != "" {
		err := u.updateUserEmail(ctx, oldUser, request.User.Email)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *userService) updateUserPassword(ctx context.Context, oldUser *model.User, oldPassword string, newPassword string) *userErrors.CustomErr {
	const op = "userService.updateUserPassword"

	err := validateUserPassword(newPassword)
	if err != nil {
		u.log.Error("error validating user password", zap.Error(err), zap.String("op", op))
		return userErrors.New(err, "invalid password", codes.InvalidArgument)
	}

	err = u.hasher.Compare(oldUser.Password, oldPassword)
	if err != nil {
		u.log.Error("error comparing password", zap.Error(err), zap.String("op", op))
		return userErrors.New(err, "invalid email", codes.InvalidArgument)
	}

	hashedPassword, err := u.hasher.Hash(newPassword)
	if err != nil {
		u.log.Error("error hashing password", zap.Error(err), zap.String("op", op))
		return userErrors.NewInternalErr(err)
	}

	updatedUser := &model.User{
		ID:        oldUser.ID,
		Name:      oldUser.Name,
		Email:     oldUser.Email,
		Password:  hashedPassword,
		Role:      oldUser.Role,
		UpdatedAt: time.Now(),
	}

	err = u.userRepo.Update(ctx, updatedUser)
	if err != nil {
		u.log.Error("error updating user", zap.Error(err), zap.String("op", op))
		return userErrors.NewInternalErr(err)
	}

	return nil
}

func (u *userService) updateUserName(ctx context.Context, oldUser *model.User, newName string) *userErrors.CustomErr {
	const op = "userService.updateUserName"

	err := validateUserName(newName)
	if err != nil {
		return userErrors.New(err, "invalid name", codes.InvalidArgument)
	}

	updatedUser := &model.User{
		ID:        oldUser.ID,
		Name:      newName,
		Email:     oldUser.Email,
		Password:  oldUser.Password,
		Role:      oldUser.Role,
		UpdatedAt: time.Now(),
	}

	err = u.userRepo.Update(ctx, updatedUser)
	if err != nil {
		return userErrors.NewInternalErr(err)
	}

	return nil
}

func (u *userService) updateUserEmail(ctx context.Context, oldUser *model.User, newEmail string) *userErrors.CustomErr {
	const op = "userService.updateUserEmail"

	valid := IsValidEmail(newEmail)
	if !valid {
		return userErrors.New(userErrors.ErrInvalidEmail, "invalid email", codes.Internal)
	}

	updatedUser := &model.User{
		ID:        oldUser.ID,
		Name:      oldUser.Name,
		Email:     newEmail,
		Password:  oldUser.Password,
		Role:      oldUser.Role,
		UpdatedAt: time.Now(),
	}

	err := u.userRepo.Update(ctx, updatedUser)
	if err != nil {
		return userErrors.NewInternalErr(err)
	}

	return nil
}

func validateUserName(name string) error {
	if name == "" || utf8.RuneCountInString(name) < 3 || utf8.RuneCountInString(name) > 30 {
		return userErrors.ErrInvalidName
	}

	return nil
}

func validateUserPassword(password string) error {
	if password == "" || utf8.RuneCountInString(password) < 8 || utf8.RuneCountInString(password) > 30 {
		return userErrors.ErrInvalidPassword
	}

	return nil
}
