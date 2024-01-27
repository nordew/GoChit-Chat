package user

import (
	"context"
	"go.uber.org/zap"
)

func (u *userService) Delete(ctx context.Context, id string) error {
	const op = "userService.Delete"

	err := u.userRepo.Delete(ctx, id)
	if err != nil {
		u.log.Error("error deleting user", zap.Error(err), zap.String("op", op))
		return err
	}

	return nil
}
