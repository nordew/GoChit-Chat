package user

import (
	"context"
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"go.uber.org/zap"
	"user/internal/api/user/dao"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	const op = "user.Implementation.Create"

	converted := dao.ToUserFromCreateRequest(req)

	id, err := i.userService.Create(ctx, converted)
	if err != nil {
		i.log.Error("error creating user", zap.Error(err), zap.String("op", op))
		return nil, err
	}

	resp := &desc.CreateUserResponse{
		Id: id,
	}

	return resp, nil
}
