package user

import (
	"context"
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"go.uber.org/zap"
	"user/internal/api/user/dao"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	const op = "user.Implementation.Get"

	i.log.Info("getting user", zap.String("op", op))
	user, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	i.log.Info("got user", zap.String("op", op))
	converted := dao.ToResponseFromUser(user)

	i.log.Info("converted user", zap.String("op", op))
	resp := &desc.GetUserResponse{
		Id:        converted.Id,
		Email:     converted.Email,
		Role:      converted.Role,
		CreatedAt: converted.CreatedAt,
		UpdatedAt: converted.UpdatedAt,
	}

	return resp, nil
}
