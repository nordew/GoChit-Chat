package user

import (
	"context"
	"user/internal/api/user/dao"

	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"go.uber.org/zap"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	const op = "user.Implementation.Create"

	converted := dao.ToUserFromCreateRequest(req)

	resp, err := i.userService.Create(ctx, converted)
	if err != nil {
		i.log.Error("error creating user", zap.Error(err), zap.String("op", op))
		return nil, err
	}

	convertedResponse := dao.ToResponseFromCreateUserResponse(resp)

	return convertedResponse, nil
}
