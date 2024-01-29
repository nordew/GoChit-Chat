package user

import (
	"context"
	"go.uber.org/zap"
	"time"
	"unicode/utf8"
	"user/internal/model"
	"user/internal/service"
	userErrors "user/pkg/user_errors"
)

func (u *userService) Update(ctx context.Context, request *service.UpdateUserRequest) error {
	const op = "userService.Update"

	oldUser, err := u.userRepo.GetById(ctx, request.User.ID)
	if err != nil {
		u.log.Error("error getting user", zap.Error(err), zap.String("op", op))
		return err
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

func (u *userService) updateUserPassword(ctx context.Context, oldUser *model.User, oldPassword string, newPassword string) error {
	const op = "userService.updateUserPassword"

	err := validateUserPassword(newPassword)
	if err != nil {
		u.log.Error("error validating user password", zap.Error(err), zap.String("op", op))
		return err
	}

	err = u.hasher.Compare(oldUser.Password, oldPassword)
	if err != nil {
		u.log.Error("error comparing password", zap.Error(err), zap.String("op", op))
		return userErrors.ErrWrongEmailOrPassword
	}

	hashedPassword, err := u.hasher.Hash(newPassword)
	if err != nil {
		u.log.Error("error hashing password", zap.Error(err), zap.String("op", op))
		return err
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
		return err
	}

	return nil
}

func (u *userService) updateUserName(ctx context.Context, oldUser *model.User, newName string) error {
	const op = "userService.updateUserName"

	err := validateUserName(newName)
	if err != nil {
		return err
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
		return err
	}

	return nil
}

func (u *userService) updateUserEmail(ctx context.Context, oldUser *model.User, newEmail string) error {
	const op = "userService.updateUserEmail"

	valid := IsValidEmail(newEmail)
	if !valid {
		return userErrors.ErrInvalidEmail
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
		return err
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
