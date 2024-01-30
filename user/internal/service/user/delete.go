package user

import (
	"context"
	userErrors "user/pkg/user_errors"

	"go.uber.org/zap"
)

func (u *userService) Delete(ctx context.Context, id string) *userErrors.CustomErr {
	const op = "userService.Delete"

	err := u.userRepo.Delete(ctx, id)
	if err != nil {
		u.log.Error("error deleting user", zap.Error(err), zap.String("op", op))
		return userErrors.NewInternalErr(err)
	}

	return nil
}
