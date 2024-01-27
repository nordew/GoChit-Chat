package user

import (
	"context"
	"go.uber.org/zap"
	"user/internal/model"
)

func (u userService) Get(ctx context.Context, email string) (*model.User, error) {
	const op = "userService.Get"

	user, err := u.userRepo.Get(ctx, email)

	if err != nil {
		u.log.Error("error getting user", zap.Error(err), zap.String("op", op))
		return nil, err
	}

	return user, nil
}
