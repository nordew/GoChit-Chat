package user

import (
	"context"
	"go.uber.org/zap"
	"time"
	"user/internal/model"
)

var ()

func (u *userService) Update(ctx context.Context, userUpdate *model.UserUpdate) error {
	const op = "userService.Update"

	err := validateUser(userUpdate.Name, userUpdate.Email, userUpdate.NewPassword)
	if err != nil {
		u.log.Error("error validating user", zap.Error(err), zap.String("op", op))
		return err
	}

	oldUser, err := u.userRepo.Get(ctx, userUpdate.ID)
	if err != nil {
		u.log.Error("error getting user", zap.Error(err), zap.String("op", op))
		return err
	}

	if userUpdate.Password {
		if err := u.changeWithPassword(ctx, userUpdate, oldUser); err != nil {
			return err
		}
	} else {
		updatedUser := getUpdatedUser(oldUser, userUpdate)
		if err := u.userRepo.Update(ctx, updatedUser); err != nil {
			u.log.Error("error updating user", zap.Error(err), zap.String("op", op))
			return err
		}
	}

	return nil
}

func getUpdatedUser(oldUser *model.User, userUpdate *model.UserUpdate) *model.User {
	updatedUser := &model.User{
		ID:        oldUser.ID,
		Name:      oldUser.Name,
		Email:     oldUser.Email,
		Password:  oldUser.Password,
		Role:      oldUser.Role,
		CreatedAt: oldUser.CreatedAt,
		UpdatedAt: time.Now(),
	}

	if userUpdate.Name != "" {
		updatedUser.Name = userUpdate.Name
	}

	if userUpdate.Email != "" {
		updatedUser.Email = userUpdate.Email
	}

	return updatedUser
}

func (u *userService) changeWithPassword(ctx context.Context, newUser *model.UserUpdate, user *model.User) error {
	const op = "userService.changeWithPassword"

	err := u.hasher.Compare(user.Password, newUser.OldPassword)
	if err != nil {
		u.log.Error("error password not matching", zap.String("op", op))
		return ErrPasswordNotMatching
	}

	hashed, err := u.hasher.Hash(newUser.NewPassword)
	if err != nil {
		u.log.Error("error hashing password", zap.Error(err), zap.String("op", op))
		return err
	}

	updatedUser := getUpdatedUser(user, newUser)
	updatedUser.Password = hashed

	if err := u.userRepo.Update(ctx, updatedUser); err != nil {
		u.log.Error("error updating user", zap.Error(err), zap.String("op", op))
		return err
	}

	return nil
}
