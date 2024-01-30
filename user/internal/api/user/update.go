package user

import (
	"context"
	"errors"
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"user/internal/api/user/dao"
	userErrors "user/pkg/user_errors"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateUserRequest) (*emptypb.Empty, error) {
	const op = "user.Implementation.Update"

	converted := dao.ToUserFromUpdateRequest(req)

	err := i.userService.Update(ctx, converted)
	if err != nil {
		var grpcCode codes.Code
		var logMessage string

		switch {
		case errors.Is(err, userErrors.ErrWrongEmailOrPassword):
			grpcCode = codes.PermissionDenied
			logMessage = "permission denied: invalid password"
		case errors.Is(err, userErrors.ErrInvalidName):
			grpcCode = codes.InvalidArgument
			logMessage = "invalid user name"
		case errors.Is(err, userErrors.ErrInvalidEmail):
			grpcCode = codes.InvalidArgument
			logMessage = "invalid email address"
		case errors.Is(err, userErrors.ErrInvalidPassword):
			grpcCode = codes.InvalidArgument
			logMessage = "invalid password"
		default:
			grpcCode = codes.Internal
			logMessage = "internal error updating user"
		}

		i.log.Error(logMessage, zap.Error(err), zap.String("op", op))
		return nil, status.Error(grpcCode, logMessage)
	}

	return &emptypb.Empty{}, nil
}
