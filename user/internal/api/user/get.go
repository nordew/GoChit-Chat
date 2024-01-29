package user

import (
	"context"
	"user/internal/api/user/dao"

	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"go.uber.org/zap"
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

func (i *Implementation) Login(ctx context.Context, req *desc.LoginUserRequest) (*desc.LoginUserResponse, error) {
	const op = "user.Implementation.Login"

	accessToken, refreshToken, err := i.userService.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		i.log.Error("error logging in user", zap.Error(err), zap.String("op", op))
		return nil, err
	}

	resp := &desc.LoginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return resp, nil
}
