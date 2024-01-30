package user

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/internal/api/user/dao"
	userErrors "user/pkg/user_errors"

	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"go.uber.org/zap"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	const op = "user.Implementation.Create"

	converted := dao.ToUserFromCreateRequest(req)

	resp, err := i.userService.Create(ctx, converted)
	if err != nil {
		var grpcCode codes.Code
		var logMessage string

		switch {
		case errors.Is(err, userErrors.ErrEmailAlreadyExists):
			grpcCode = codes.AlreadyExists
			logMessage = "user with this email already exists"
		default:
			grpcCode = codes.Internal
			logMessage = "internal error creating user"
		}

		i.log.Error(logMessage, zap.Error(err), zap.String("op", op))
		return nil, status.Error(grpcCode, logMessage)
	}

	convertedResponse := dao.ToResponseFromCreateUserResponse(resp)

	return convertedResponse, nil
}
