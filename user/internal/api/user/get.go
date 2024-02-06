package user

import (
	"context"
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"user/internal/api/user/dao"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	const op = "user.Implementation.Get"

	i.log.Info("getting user", zap.String("op", op))
	user, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		i.log.Error("error getting user", zap.Error(err.Err), zap.String("op", op))
		return nil, status.Error(err.Code, err.Msg)
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
		i.log.Error("error logging in user", zap.Error(err.Err), zap.String("op", op))
		return nil, status.Error(err.Code, err.Msg)
	}

	resp := &desc.LoginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return resp, nil
}

func (i *Implementation) ParseAccessToken(ctx context.Context, req *desc.ParseAccessTokenRequest) (*desc.ParseAccessTokenResponse, error) {
	const op = "user.Implementation.ParseAccessToken"

	userID, name, err := i.userService.ParseAccessToken(ctx, req.Token)
	if err != nil {
		i.log.Error("error parsing token", zap.Error(err.Err), zap.String("op", op))
		return nil, status.Error(err.Code, err.Msg)
	}

	resp := &desc.ParseAccessTokenResponse{
		UserID: userID,
		Name:   name,
	}

	return resp, nil
}

func (i *Implementation) Refresh(ctx context.Context, req *desc.RefreshUserRequest) (*desc.RefreshUserResponse, error) {
	const op = "user.Implementation.Refresh"

	accessToken, refreshToken, err := i.userService.Refresh(ctx, req.GetRefreshToken())
	if err != nil {
		i.log.Error("error refreshing tokens", zap.Error(err.Err), zap.String("op", op))
		return nil, status.Error(err.Code, err.Msg)
	}

	resp := &desc.RefreshUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return resp, nil
}
