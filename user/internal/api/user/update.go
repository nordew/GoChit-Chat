package user

import (
	"context"
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"user/internal/api/user/dao"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateUserRequest) (*emptypb.Empty, error) {
	const op = "user.Implementation.Update"

	converted := dao.ToUserFromUpdateRequest(req)

	err := i.userService.Update(ctx, converted)
	if err != nil {
		i.log.Error("failed to update user", zap.Error(err.Err), zap.String("op", op))
		return nil, status.Error(err.Code, err.Msg)
	}

	return &emptypb.Empty{}, nil
}
