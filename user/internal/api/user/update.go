package user

import (
	"context"
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"user/internal/api/user/dao"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateUserRequest) (*emptypb.Empty, error) {
	const op = "user.Implementation.Update"

	converted := dao.ToUserFromUpdateRequest(req)

	err := i.userService.Update(ctx, converted)
	if err != nil {
		i.log.Error("error updating user", zap.Error(err), zap.String("op", op))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
