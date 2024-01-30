package user

import (
	"context"
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"google.golang.org/grpc/status"
	"user/internal/api/user/dao"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	const op = "user.Implementation.Create"

	converted := dao.ToUserFromCreateRequest(req)

	resp, err := i.userService.Create(ctx, converted)
	if err != nil {
		return nil, status.Error(err.Code, err.Msg)
	}

	convertedResponse := dao.ToResponseFromCreateUserResponse(resp)

	return convertedResponse, nil
}
