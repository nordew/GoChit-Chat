package user

import (
	"context"
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteUserRequest) (*emptypb.Empty, error) {
	const op = "user.Implementation.Delete"

	err := i.userService.Delete(ctx, req.GetId())
	if err != nil {
		i.log.Error("failed to delete user", zap.Error(err.Err), zap.String("op", op))
		return nil, status.Error(err.Code, err.Msg)
	}

	return &emptypb.Empty{}, nil
}
