package user

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/internal/api/user/dao"
	userErrors "user/pkg/user_errors"

	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"go.uber.org/zap"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	const op = "user.Implementation.Get"

	i.log.Info("getting user", zap.String("op", op))
	user, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		var grpcCode codes.Code
		var logMessage string

		switch {
		case errors.Is(err, userErrors.ErrWrongEmailOrPassword):
			grpcCode = codes.NotFound
			logMessage = "user not found"
		default:
			grpcCode = codes.Internal
			logMessage = "internal error getting user"
		}

		i.log.Error(logMessage, zap.Error(err), zap.String("op", op))
		return nil, status.Error(grpcCode, logMessage)
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
		var grpcCode codes.Code
		var logMessage string

		switch {
		case errors.Is(err, userErrors.ErrInvalidEmail):
			grpcCode = codes.InvalidArgument
			logMessage = "invalid email"
		case errors.Is(err, userErrors.ErrWrongEmailOrPassword):
			grpcCode = codes.InvalidArgument
			logMessage = "wrong email or password"
		default:
			grpcCode = codes.Internal
			logMessage = "internal error while logging in"
		}

		i.log.Error("error logging in user", zap.Error(err), zap.String("op", op))
		return nil, status.Error(grpcCode, logMessage)
	}

	resp := &desc.LoginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return resp, nil
}
